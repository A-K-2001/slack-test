package linear

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
)

var apiKey string

func CreateLinearClient(ApiKey string) {
	apiKey = ApiKey
}

func CreateTriage(teamID, title, description string) (string, error) {
	url := "https://api.linear.app/graphql"

	query := `
        mutation CreateIssue($input: IssueCreateInput!) {
            issueCreate(input: $input) {
                success
                issue {
                    id
                    title
                }
            }
        }
    `

	input := map[string]interface{}{
		"input": map[string]interface{}{
			"teamId":      teamID,
			"title":       title,
			"description": description,
		},
	}

	payload := map[string]interface{}{
		"query":     query,
		"variables": input,
	}

	payloadBytes, _ := json.Marshal(payload)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status: %s, response body: %s", resp.Status, string(body))
	}

	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return "", err
	}

	issueCreate, ok := result["data"].(map[string]interface{})["issueCreate"].(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("unexpected response format: %v", result)
	}

	issue, ok := issueCreate["issue"].(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("unexpected response format: %v", issueCreate)
	}

	issueID, ok := issue["id"].(string)
	if !ok {
		return "", fmt.Errorf("issue ID not found in response: %v", issue)
	}

	return issueID, nil
}

func UploadScreenshotToIssue(issueID, fileURL, fileName string) error {
	url := "https://api.linear.app/graphql"

	// Prepare the mutation query
	query := `
        mutation CreateAttachment($input: AttachmentCreateInput!) {
            attachmentCreate(input: $input) {
                success
                attachment {
                    id
                    url
                }
            }
        }
    `

	// Prepare the variables for the mutation
	variables := map[string]interface{}{
		"input": map[string]interface{}{
			"issueId": issueID,
			"title":   fileName, // Title for the attachment
			"url":     fileURL,  // Direct URL to the uploaded image
		},
	}

	payload := map[string]interface{}{
		"query":     query,
		"variables": variables,
	}

	payloadBytes, _ := json.Marshal(payload)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Authorization", apiKey) // No Bearer prefix
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to execute request: %v", err)
	}
	defer resp.Body.Close()

	responseBody, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status: %s, response body: %s", resp.Status, string(responseBody))
	}

	fmt.Println("Screenshot uploaded successfully")
	return nil
}

func GetIssueId(ev *slackevents.FileSharedEvent, client *slack.Client) string {

	fileInfo, _, _, err := client.GetFileInfo(ev.FileID, 0, 0)
	if err != nil {
		log.Printf("Error fetching file info: %v", err)
		return "Error fetching file info"
	}

	timestamp := ""

	for _, fileInfos := range fileInfo.Shares.Private {
		if len(fileInfos) > 0 {
			threadTs := fileInfos[0].ThreadTs
			// fmt.Printf("Key: %s, ThreadTs from Info 0: %s\n", key, threadTs)
			timestamp = threadTs
			break
		}
	}
	print("threadts", timestamp)
	params := slack.GetConversationHistoryParameters{
		ChannelID:          ev.ChannelID,
		Latest:             timestamp,
		Inclusive:          true,
		Limit:              1,
		IncludeAllMetadata: true,
	}

	history, err := client.GetConversationHistory(&params)
	if err != nil {
		log.Fatalf("Error fetching conversation history: %v", err)
	}

	if len(history.Messages) > 0 {
		message := history.Messages[0]
		issueId := message.Msg.Metadata.EventPayload["IssueId"]
		return issueId.(string)
	} else {
		return "No message found at that timestamp."
	}
}
