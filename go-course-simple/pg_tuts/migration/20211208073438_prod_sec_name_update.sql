-- +goose Up
-- +goose StatementBegin
ALTER TABLE product_items_collection ADD COLUMN prod_sec_name text NOT NULL DEFAULT '';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE product_items_collection DROP COLUMN prod_sec_name;
-- +goose StatementEnd
