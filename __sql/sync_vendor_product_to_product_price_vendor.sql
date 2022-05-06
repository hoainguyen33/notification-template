--region vendor product change

-- INSERT getcare_product_price_vendor and flag log if exist
insert into getcare_product_price_vendor(getcare_product_id, getcare_vendor_id, price_buy, price_sales, price_sales_retail, vat, operation, amount, getcare_policy_unit_id, getcare_uom_base_id, getcare_user_id, updated_at_price_buy, updated_at_price_sales, updated_at_price_sales_retail, minimum_quantity, estimated_quantity)
select getcare_product_id, getcare_vendor_id, price_buy, price_sales, price_sales_retail, vat, operation, amount, getcare_policy_unit_id, getcare_uom_base_id, getcare_user_id, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, minimum_quantity, estimated_quantity
from getcare_vendor_product
where getcare_product_id is not null and getcare_uom_base_id is not null and deleted=0 and getcare_vendor_id=?
    on duplicate key update log=IF(
        getcare_product_price_vendor.price_buy <> getcare_vendor_product.price_buy
        or getcare_product_price_vendor.price_sales <> getcare_vendor_product.price_sales
        or IFNULL(getcare_product_price_vendor.price_sales_retail, 0) <> IFNULL(getcare_vendor_product.price_sales_retail, 0)
        or IFNULL(getcare_product_price_vendor.vat,-1) <> IFNULL(getcare_vendor_product.vat,-1)
        or IFNULL(getcare_product_price_vendor.estimated_quantity,-1) <> IFNULL(getcare_vendor_product.estimated_quantity,-1)
        or IFNULL(getcare_product_price_vendor.minimum_quantity,-1) <> IFNULL(getcare_vendor_product.minimum_quantity,-1), 1, log);
-- FLAG delete getcare_product_price_vendor
update getcare_product_price_vendor gppv
    left join getcare_vendor_product gp
on gppv.getcare_vendor_id=gp.getcare_vendor_id
    and gppv.getcare_product_id=gp.getcare_product_id
    and gppv.getcare_uom_base_id=gp.getcare_uom_base_id and gp.deleted=0
set gppv.`delete`=1
    where gppv.getcare_vendor_id=? and gp.id is null;

-- LOG getcare_product_price_vendor_history
insert into getcare_product_price_vendor_history(getcare_product_price_vendor_id, getcare_product_id, getcare_vendor_id, price_buy, price_sales, price_buy_recently, price_sales_recently, getcare_uom_base_id, getcare_user_id, channel_source, updated_at, updated_at_price_buy, updated_at_price_sales, updated_at_price_buy_recently, updated_at_price_sales_recently, vat)
select id, getcare_product_id, getcare_vendor_id, price_buy, price_sales, price_buy_recently, price_sales_recently, getcare_uom_base_id, getcare_user_id, channel_source, updated_at, updated_at_price_buy, updated_at_price_sales, updated_at_price_buy_recently, updated_at_price_sales_recently, vat
from getcare_product_price_vendor where log=1 or `delete`=1;

-- UPDATE getcare_product_price_vendor WHERE flag log
update getcare_product_price_vendor gppv left join getcare_vendor_product gp on
    gppv.getcare_vendor_id = gp.getcare_vendor_id and gppv.getcare_product_id = gp.getcare_product_id and
    gppv.getcare_uom_base_id = gp.getcare_uom_base_id and gp.deleted = 0
set gppv.price_buy_recently=IF(IFNULL(gppv.price_buy,-1)=IFNULL(gp.price_buy,-1), gppv.price_buy_recently, gppv.price_buy),
    gppv.price_buy=gp.price_buy,
    gppv.updated_at_price_buy_recently=IF(IFNULL(gppv.price_buy,-1)=IFNULL(gp.price_buy,-1), gppv.updated_at_price_buy_recently, gppv.updated_at_price_buy),
    gppv.updated_at_price_buy=IF(IFNULL(gppv.price_buy,-1)=IFNULL(gp.price_buy,-1), gppv.updated_at_price_buy, CURRENT_TIMESTAMP),

    gppv.price_sales_recently=IF(IFNULL(gppv.price_sales,-1)=IFNULL(gp.price_sales,-1), gppv.price_sales_recently, gppv.price_sales),
    gppv.price_sales=gp.price_sales,
    gppv.updated_at_price_sales_recently=IF(IFNULL(gppv.price_sales,-1)=IFNULL(gp.price_sales,-1), gppv.updated_at_price_sales_recently, gppv.updated_at_price_sales),
    gppv.updated_at_price_sales=IF(IFNULL(gppv.price_sales,-1)=IFNULL(gp.price_sales,-1), gppv.updated_at_price_sales, CURRENT_TIMESTAMP),

    gppv.price_sales_retail_recently=IF(IFNULL(gppv.price_sales_retail,-1)=IFNULL(gp.price_sales_retail,-1), gppv.price_sales_retail_recently, gp.price_sales_retail),
    gppv.price_sales_retail=gp.price_sales_retail,
    gppv.updated_at_price_sales_retail_recently=IF(IFNULL(gppv.price_sales_retail,-1)=IFNULL(gp.price_sales_retail,-1), gppv.updated_at_price_sales_retail_recently, gppv.updated_at_price_sales_retail),
    gppv.updated_at_price_sales_retail=IF(IFNULL(gppv.price_sales_retail,-1)=IFNULL(gp.price_sales_retail,-1), gppv.updated_at_price_sales_retail, CURRENT_TIMESTAMP),

    gppv.vat=gp.vat,
    gppv.minimum_quantity=gp.minimum_quantity,
    gppv.remaining_quantity=IF(IFNULL(gppv.estimated_quantity,-1)=IFNULL(gp.estimated_quantity,-1), gppv.remaining_quantity, gp.estimated_quantity),
    gppv.estimated_quantity=gp.estimated_quantity,
    gppv.operation=gp.operation,
    gppv.amount=gp.amount,
    gppv.getcare_policy_unit_id=gp.getcare_policy_unit_id,
    log=0
where log = 1;

-- RESET flag log
update getcare_product_price_vendor set log=0 where log=1;

-- DELETE record flag delete
delete from getcare_product_price_vendor where `delete`=1;

-- SYNC type_label
UPDATE getcare_product_price_vendor gppv
    JOIN getcare_vendor_product gvp ON gppv.getcare_vendor_id=gvp.getcare_vendor_id AND gppv.getcare_product_id=gvp.getcare_product_id AND gppv.getcare_uom_base_id=gvp.getcare_uom_base_id
SET gppv.type_label=gvp.type_label
WHERE (gppv.type_label IS NULL AND gvp.type_label IS NOT NULL) OR (gppv.type_label IS NOT NULL AND gvp.type_label IS NULL) OR gppv.type_label != gvp.type_label

--endregion