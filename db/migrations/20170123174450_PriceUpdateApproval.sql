-- +goose Up
CREATE TABLE PriceUpdateApproval (
  product_id INT  NOT NULL,
  version    TEXT NOT NULL,
  status     TEXT NOT NULL
);
-- SQL in section 'Up' is executed when this migration is applied


-- +goose Down
DROP TABLE PriceUpdateApproval;
-- SQL section 'Down' is executed when this migration is rolled back

