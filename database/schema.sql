CREATE TYPE rider_status AS ENUM ('online', 'offline');

CREATE TYPE vehicle_type AS ENUM ('bike', 'truck', 'car');

CREATE TYPE trip_type AS ENUM ('shared', 'forward', 'return');

CREATE TYPE trip_status AS ENUM (
    'new',
    'rider_assigned',
    'enroute_to_pickup',
    'arrived_at_dropoff',
    'enroute_to_dropoff',
    'nearby_dropoff',
    'completed',
    'cancelled'
);

CREATE TYPE trip_checkpoint_type AS ENUM (
    'pickup',
    'dropoff',
    'rerouted',
    'intermediate',
    'dropped'
);

CREATE TYPE logistic_event_type AS ENUM (
    'rerouted',
    'driver_update',
    'arrived',
    'visited',
    'assigned',
    'picked',
    'dropped'
);

CREATE TABLE "rider" (
    "user_id" text UNIQUE PRIMARY KEY,
    "status" rider_status NOT NULL,
    "geohash" text NOT NULL,
    "battery" int DEFAULT 0 NOT NULL,
    "vehicle_type" vehicle_type NOT NULL,
    "created_at" timestamptz DEFAULT NOW(),
    "updated_at" timestamptz DEFAULT NOW()
);

CREATE TABLE "trip" (
    "trip_id" uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    "external_trip_id" text,
    "encoded_polyline" text,
    "rider_id" text,
    "trip_type" trip_type NOT NULL,
    "dropoff_geohash" text NOT NULL,
    "pickup_geohash" text NOT NULL,
    "status" trip_status NOT NULL DEFAULT 'new',
    "created_at" timestamptz DEFAULT NOW(),
    "updated_at" timestamptz DEFAULT NOW()
);

CREATE TABLE "trip_checkpoint" (
    "id" uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    "trip_id" uuid NOT NULL,
    "checkpoint_order" int NOT NULL,
    "checkpoint_type" trip_checkpoint_type NOT NULL,
    "arrived_event" uuid,
    "visited_event" uuid,
    "checkpoint_geohash" text NOT NULL,
    "aggregated_trip_distance" int NOT NULL,
    "aggregated_estimated_time" int NOT NULL,
    "created_at" timestamptz DEFAULT NOW(),
    "updated_at" timestamptz DEFAULT NOW()
);

CREATE TABLE "logistic_events" (
    "id" uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    "rider_id" text,
    "event_type" logistic_event_type NOT NULL,
    "event_battery_percentage" int NOT NULL,
    "event_geohash" text NOT NULL,
    "event_payload" json DEFAULT '{}',
    "created_at" timestamptz DEFAULT NOW(),
    "updated_at" timestamptz DEFAULT NOW()
);

ALTER TABLE
    "trip"
ADD
    FOREIGN KEY ("rider_id") REFERENCES "rider" ("user_id");

ALTER TABLE
    "trip_checkpoint"
ADD
    FOREIGN KEY ("trip_id") REFERENCES "trip" ("trip_id");

-- CREATE OR REPLACE FUNCTION updated_at_column_procedure()
--     RETURNS TRIGGER AS
-- $$
-- BEGIN
--     NEW.updated_at = now();
-- RETURN NEW;
-- END;
-- $$ language 'plpgsql';

CREATE TRIGGER change_updated_at_on_logistic_events BEFORE UPDATE
    ON logistic_events FOR EACH ROW EXECUTE PROCEDURE
    updated_at_column_procedure();

CREATE TRIGGER change_updated_at_on_rider BEFORE UPDATE
    ON rider FOR EACH ROW EXECUTE PROCEDURE
    updated_at_column_procedure();

CREATE TRIGGER change_updated_at_on_trip BEFORE UPDATE
    ON trip FOR EACH ROW EXECUTE PROCEDURE
    updated_at_column_procedure();

CREATE TRIGGER change_updated_at_on_trip_checkpoint BEFORE UPDATE
    ON trip_checkpoint FOR EACH ROW EXECUTE PROCEDURE
    updated_at_column_procedure();

-- HSN schema

CREATE TYPE Platform AS ENUM ('ANDROID', 'IOS', 'WEB', 'SERVICE', 'UNKNOWN');
CREATE TYPE MarketPlace AS ENUM ('PLAY_STORE', 'APP_STORE');


CREATE TYPE Role AS ENUM ('DRIVER', 'CASHIER', 'SUPPORT');
CREATE TYPE PrescriptionStatus AS ENUM ('APPROVED', 'REJECTED', 'PENDING', 'DELETED');
CREATE TYPE HospitalRoles AS ENUM ('DOCTOR', 'RECEPTIONIST', 'TECHNICIAN');
CREATE TYPE SaleChannel AS ENUM ('ONLINE', 'OFFLINE', 'OMNI', 'NOT_FOR_SALE');
CREATE TYPE MedProcurementStatus AS ENUM ('PURCHASED', 'ORDERED', 'NOT_IN_DEMAND', 'NOT_IN_MARKET');
CREATE TYPE DiscountTier AS ENUM (
  'AYURVEDIC', 'COSMETICS', 'EQUIPMENTS', 'INJECTIONS', 'MEDICINES', 'OTC', 'SURGICALS',
  'DELIVERY_CHARGE', 'LATE_NIGHT_FEE', 'SMALL_CART_FEE', 'AVD_ABOVE_30', 'OTC_BELOW_10',
  'COS_ABOVE_40', 'EQP_BELOW_20', 'SGL_BELOW_20', 'COS_BELOW_40', 'EQP_BELOW_30',
  'OTC_BELOW_20', 'OTC_ABOVE_30', 'SGL_ABOVE_40', 'COS_BELOW_30', 'OTC_BELOW_15',
  'OTC_BELOW_30', 'INJ_BELOW_25', 'SGL_BELOW_30', 'MED_BELOW_30', 'INJ_BELOW_20',
  'MED_BELOW_50', 'OTC_BELOW_25', 'MED_BELOW_20', 'MED_ABOVE_50', 'SGL_BELOW_40',
  'MED_BELOW_10', 'MED_BELOW_40', 'INJ_ABOVE_30', 'OTC_BELOW_5', 'EQP_ABOVE_40',
  'EQP_BELOW_40', 'AVD_BELOW_30', 'INJ_BELOW_30', 'AVD_BELOW_20', 'INJ_BELOW_15',
  'COS_BELOW_20'
);
CREATE TYPE DrugSchedule AS ENUM (
  'A', 'B', 'C', 'C1', 'D', 'E', 'E1', 'F', 'F1', 'F2', 'F3', 'FF', 'G', 'H', 'J',
  'K', 'M', 'M1', 'M2', 'M3', 'N', 'O', 'P', 'Q', 'R', 'R1', 'S', 'T', 'U', 'U1',
  'V', 'W', 'X', 'Y', 'H1'
);
CREATE TYPE MedDemandLedgerStatus AS ENUM ('DRAFT', 'FULFILLED', 'IGNORED');
CREATE TYPE MedDemandLedgerType AS ENUM ('AGENT');


CREATE TABLE "app_updates" (
  "id" serial PRIMARY KEY,
  "product" text ,
  "platform" platform NOT NULL ,
  "current_build" integer,
  "deprecated_build" integer ,
  "discountinued_build" integer ,
  "created_at" timestamptz  DEFAULT now(),
  "updated_at" timestamptz  DEFAULT now(),
  "marketplace" marketplace ,
  UNIQUE ("product", "platform")
);

CREATE TABLE "meds_pharmacy" (
  "id" uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  "name" varchar(100) NOT NULL,
  "latitude" decimal(10, 8) NOT NULL,
  "longitude" decimal(11, 8) NOT NULL,
  "address" text NOT NULL,
  "phone" varchar(15),
  "is_active" boolean NOT NULL DEFAULT false,
  "printer_id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "dl_number" varchar(50),
  "fssai_number" varchar(50),
  "gst_number" varchar(15),
  "pan_number" varchar(10),
  "shipping_partner" text NOT NULL DEFAULT 'FARMAKO',
  "pref_otp_enabled_delivery" boolean NOT NULL DEFAULT false,
  "dummy_cart_id" text
);


CREATE TABLE "med_order_attributes" (
  "id" serial PRIMARY KEY,
  "order_id" varchar(20) UNIQUE NOT NULL,
  "user_id" varchar(36) NOT NULL,
  "app_info" jsonb DEFAULT '{}',
  "balance" integer NOT NULL DEFAULT 0,
  "is_paid" boolean NOT NULL DEFAULT false,
  "bill_status" varchar(10),
  "ce_channel_id" integer,
  "payment_mode" varchar,
  "payment_details" jsonb DEFAULT '{}',
  "customer_address" varchar,
  "customer_location" jsonb DEFAULT '{}',
  "discount_percentage" integer,
  "pharmacy_delivery_eta" integer,
  "customer_pharmacy_distance" integer,
  "is_agent_order" boolean NOT NULL DEFAULT false,
  "pg_plink_id" varchar,
  "logistic_details" jsonb DEFAULT '{}',
  "business_id" varchar,
  "discount_amount" integer,
  "created_at" timestamptz DEFAULT now(),
  "serving_pharmacy_id" uuid,
  "billing_name" text,
  "prescription_note" text NOT NULL DEFAULT '',
  "total_amount_paisa" integer,
  FOREIGN KEY ("serving_pharmacy_id") REFERENCES "meds_pharmacy" ("id")
);

CREATE TABLE "logistic_trips" (
  "trip_id" uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  "order_id" varchar(20) UNIQUE NOT NULL,
  "status" text NOT NULL DEFAULT 'NEW',
  "partner" varchar(30) NOT NULL,
  "partner_payload" jsonb NOT NULL DEFAULT '{}',
  "enroute_payload" jsonb NOT NULL DEFAULT '{}',
  "created_at" timestamptz NOT NULL DEFAULT now(),
  "updated_at" timestamptz NOT NULL DEFAULT now(),
  "vehicle_id" text,
  "cash_collected" integer NOT NULL DEFAULT 0,
  FOREIGN KEY ("order_id") REFERENCES "med_order_attributes" ("order_id")
);


CREATE TABLE "meds_savedAddress" (
  "id" serial PRIMARY KEY,
  "label" varchar(100),
  "user_id" varchar(36) NOT NULL,
  "address_line_1" varchar(200) NOT NULL,
  "city" varchar(50) NOT NULL,
  "pin_code" varchar(20) NOT NULL,
  "address_street" varchar(100) NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT now(),
  "latitude" float NOT NULL,
  "longitude" float NOT NULL,
  UNIQUE ("user_id", "label")
);

CREATE TABLE "cart" (
  "cart_id" uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  "order_id" varchar(20) UNIQUE NOT NULL,
  "uniquer_floor" text,
  "uniquer_round" text,
  "address_id" integer,
  "cart_status" text NOT NULL DEFAULT 'DRAFT',
  "verification_status" text NOT NULL DEFAULT 'VERIFIED',
  "agent_id" text,
  "created_at" timestamptz NOT NULL DEFAULT now(),
  "billed_at" bigint,
  FOREIGN KEY ("address_id") REFERENCES "meds_savedAddress" ("id"),
  FOREIGN KEY ("order_id") REFERENCES "med_order_attributes" ("order_id"),
  UNIQUE ("cart_id", "order_id"),
  UNIQUE ("uniquer_floor", "address_id"),
  UNIQUE ("uniquer_round", "address_id")
);


CREATE TABLE "cart_item" (
  "id" uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  "cart_id" uuid NOT NULL,
  "item_type" text NOT NULL,
  "item_name" text NOT NULL,
  "discount" integer NOT NULL DEFAULT 0,
  "version" integer NOT NULL DEFAULT 1,
  "quantity" integer NOT NULL,
  "inventory_id" uuid,
  "amount" integer NOT NULL DEFAULT 0,
  "cgst" integer NOT NULL DEFAULT 0,
  "igst" integer NOT NULL DEFAULT 0,
  "sgst" integer NOT NULL DEFAULT 0,
  FOREIGN KEY ("cart_id") REFERENCES "cart" ("cart_id"),
  UNIQUE ("inventory_id", "cart_id")
);


CREATE TABLE "med_details" (
  "id" uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  "name" text NOT NULL,
  "salt" text,
  "manufacturer" text NOT NULL,
  "size_string" text NOT NULL,
  "size" integer NOT NULL,
  "default_discount" integer NOT NULL,
  "prescription_needed" boolean NOT NULL,
  "hsn" text NOT NULL,
  "max_orderable_qty" integer NOT NULL DEFAULT 100,
  "medicine_type" text NOT NULL DEFAULT 'DRUG',
  "description" text NOT NULL DEFAULT '',
  "is_discontinued" boolean NOT NULL DEFAULT false,
  "mrp_pack" integer NOT NULL DEFAULT 0,
  "slug_serial" serial NOT NULL,
  "discount_tier" discount_tier NOT NULL,
  "sale_channel" sale_channel NOT NULL DEFAULT 'OMNI',
  "is_old" boolean NOT NULL DEFAULT false,
  "external_slug" text
);


CREATE TABLE "med_batch" (
  "id" uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  "medicine_id" uuid NOT NULL,
  "batch" text NOT NULL,
  "expiry" text NOT NULL,
  "expiry_date" text NOT NULL,
  "gst_slab" integer NOT NULL,
  "ptr_pack" integer NOT NULL,
  "mrp_pack" integer NOT NULL,
  FOREIGN KEY ("medicine_id") REFERENCES "med_details" ("id") ON DELETE CASCADE,
  UNIQUE ("batch", "medicine_id"),
  UNIQUE ("medicine_id", "batch")
);


CREATE TABLE "med_stock" (
  "id" serial PRIMARY KEY,
  "pharmacy_id" uuid NOT NULL,
  "medicine_id" uuid NOT NULL,
  "batch" text NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT now(),
  "stock" integer NOT NULL DEFAULT 0,
  FOREIGN KEY ("medicine_id", "batch") REFERENCES "med_batch" ("medicine_id", "batch") ON DELETE CASCADE,
  UNIQUE ("pharmacy_id", "medicine_id", "batch")
);

