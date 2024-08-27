--
-- Copyright (C) 2024 IOTech Ltd
--
-- SPDX-License-Identifier: Apache-2.0

-- app_svc.store is used to save StoredObject as JSONB on failure
CREATE TABLE IF NOT EXISTS app_svc.store (
    id UUID PRIMARY KEY,
    content JSONB NOT NULL
);
