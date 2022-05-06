--region OVERRIDE
-- LOG getcare_vendor_product_history
INSERT INTO getcare_vendor_product_history(getcare_vendor_product_id, getcare_vendor_id, product_vendor_code, product_vendor_name, getcare_product_id, price_buy, getcare_uom_base_id, getcare_user_id, price_sales, vat, operation, amount, getcare_policy_unit_id, type_label)
SELECT id as getcare_vendor_product_id, getcare_vendor_id, product_vendor_code, product_vendor_name, getcare_product_id, price_buy, getcare_uom_base_id, getcare_user_id, price_sales, vat, operation, amount, getcare_policy_unit_id, type_label
FROM getcare_vendor_product
WHERE deleted=0 AND getcare_vendor_id=1;
-- END LOG

-- FLAG delete all
UPDATE getcare_vendor_product SET deleted=1 WHERE getcare_vendor_id=%d AND deleted=0;

-- DELETE CONFLICT getcare_vendor_product
UPDATE getcare_vendor_product gvp
JOIN getcare_vendor_product_import_item ii
ON ii.getcare_vendor_id=gvp.getcare_vendor_id
AND ii.product_vendor_code=gvp.product_vendor_code
AND ii.getcare_vendor_product_import_id=170
AND ii.getcare_uom_base_id=gvp.getcare_uom_base_id
AND ii.getcare_product_id IS NOT NULL
AND gvp.getcare_product_id IS NOT NULL
SET gvp.getcare_product_id=NULL
WHERE gvp.getcare_product_id <> ii.getcare_product_id;

-- DELETE CONFLICT getcare_vendor_product
UPDATE getcare_vendor_product gvp
JOIN getcare_vendor_product_import_item ii
ON ii.getcare_product_id=gvp.getcare_product_id
AND ii.getcare_vendor_id=gvp.getcare_vendor_id
AND ii.getcare_vendor_product_import_id=170
AND ii.getcare_uom_base_id=gvp.getcare_uom_base_id
AND ii.product_vendor_code IS NOT NULL
AND gvp.product_vendor_code IS NOT NULL
SET gvp.getcare_product_id=NULL
WHERE gvp.product_vendor_code <> ii.product_vendor_code;

-- INSERT all record from getcare_vendor_product_import_item to getcare_vendor_product
INSERT INTO getcare_vendor_product(getcare_vendor_id, product_vendor_code, product_vendor_name, getcare_product_id, price_buy, getcare_uom_base_id, getcare_user_id, price_sales, vat, operation, amount, getcare_policy_unit_id, type_label, estimated_quantity, raw_data)
SELECT getcare_vendor_id, product_vendor_code, product_vendor_name, getcare_product_id, price_buy, getcare_uom_base_id, getcare_user_id, price_sales, vat, operation, amount, getcare_policy_unit_id, type_label, estimated_quantity, raw_data
FROM getcare_vendor_product_import_item gimport
WHERE getcare_vendor_product_import_id=%d
    ON DUPLICATE KEY UPDATE
        product_vendor_code=gimport.product_vendor_code,
        product_vendor_name=gimport.product_vendor_name,
        getcare_product_id=gimport.getcare_product_id,
        price_buy=gimport.price_buy,
        getcare_uom_base_id=gimport.getcare_uom_base_id,
        getcare_user_id=gimport.getcare_user_id,
        price_sales=gimport.price_sales,
        vat=gimport.vat,
        operation=gimport.operation,
        amount=gimport.amount,
        getcare_policy_unit_id=gimport.getcare_policy_unit_id,
        type_label=gimport.type_label,
        estimated_quantity=gimport.estimated_quantity,
        raw_data=gimport.raw_data,
        deleted=0;

--endregion

--region ADD
-- LOG getcare_vendor_product_history
INSERT INTO getcare_vendor_product_history(getcare_vendor_product_id, getcare_vendor_id, product_vendor_code, product_vendor_name, getcare_product_id, price_buy, getcare_uom_base_id, getcare_user_id, price_sales, vat, operation, amount, getcare_policy_unit_id, type_label)
SELECT gp.id, gp.getcare_vendor_id, gp.product_vendor_code, gp.product_vendor_name, gp.getcare_product_id, gp.price_buy, gp.getcare_uom_base_id, gp.getcare_user_id, gp.price_sales, gp.vat, gp.operation, gp.amount, gp.getcare_policy_unit_id, gp.type_label
FROM getcare_vendor_product gp
         LEFT JOIN getcare_vendor_product_import_item gi ON gp.getcare_vendor_id=gi.getcare_vendor_id AND gp.product_vendor_code=gi.product_vendor_code AND gp.deleted=0
WHERE gi.getcare_vendor_product_import_id=%d AND gi.id IS NOT NULL;

-- INSERT all record from getcare_vendor_product_import_item to getcare_vendor_product
INSERT INTO getcare_vendor_product(getcare_vendor_id, product_vendor_code, product_vendor_name, getcare_product_id, price_buy, getcare_uom_base_id, getcare_user_id, price_sales, vat, operation, amount, getcare_policy_unit_id, type_label, estimated_quantity, raw_data)
SELECT getcare_vendor_id, product_vendor_code, product_vendor_name, getcare_product_id, price_buy, getcare_uom_base_id, getcare_user_id, price_sales, vat, operation, amount, getcare_policy_unit_id, type_label, estimated_quantity, raw_data
FROM getcare_vendor_product_import_item gimport
WHERE getcare_vendor_product_import_id=%d
    ON DUPLICATE KEY UPDATE
        product_vendor_code=gimport.product_vendor_code,
        product_vendor_name=gimport.product_vendor_name,
        getcare_product_id=gimport.getcare_product_id,
        price_buy=gimport.price_buy,
        getcare_uom_base_id=gimport.getcare_uom_base_id,
        getcare_user_id=gimport.getcare_user_id,
        price_sales=gimport.price_sales,
        vat=gimport.vat,
        operation=gimport.operation,
        amount=gimport.amount,
        getcare_policy_unit_id=gimport.getcare_policy_unit_id,
        type_label=gimport.type_label,
        estimated_quantity=gimport.estimated_quantity,
        raw_data=gimport.raw_data,
        deleted=0
--endregion