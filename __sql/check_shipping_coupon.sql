SET @customer_id=136, @vendor_id=13, @price_check=410000, @date_check='2021-04-07';
SELECT getcare_shipping_coupon.* FROM getcare_shipping_coupon
  LEFT JOIN (SELECT getcare_shipping_coupon_customer.* FROM getcare_shipping_coupon_customer JOIN getcare_erp_group geg ON geg.id = getcare_erp_group_id AND geg.active=1) gscc ON getcare_shipping_coupon.id=gscc.getcare_shipping_coupon_id # group config shipping coupon
  LEFT JOIN getcare_shipping_coupon_sales_channel gscsc ON getcare_shipping_coupon.id=gscsc.getcare_shipping_coupon_id # channel config shipping coupon

  LEFT JOIN (SELECT * FROM getcare_erp_group_customer WHERE getcare_customer_id=@customer_id) gegc ON gegc.getcare_erp_group_id=gscc.getcare_erp_group_id # group khach hang = group config shipping coupon
  LEFT JOIN getcare_erp_group_customer_channel gegcc ON gegcc.getcare_erp_group_id=gscc.getcare_erp_group_id # channel cua group config shipping coupon

  LEFT JOIN getcare_sales_channel_item gsci ON gsci.getcare_sales_channel_id=gscsc.getcare_sales_channel_id OR gsci.getcare_sales_channel_id=gegcc.getcare_sales_channel_id # item cua channel
  LEFT JOIN (SELECT * from getcare_customer where id=@customer_id) gc ON gc.id = gegc.getcare_customer_id OR (gsci.getcare_ward_id=IFNULL(gc.getcare_ward_id, 0) OR (gsci.getcare_ward_id IS NULL AND gsci.getcare_district_id=IFNULL(gc.getcare_district_id, 0)) OR (gsci.getcare_ward_id IS NULL AND gsci.getcare_district_id IS NULL AND gsci.getcare_province_id=IFNULL(gc.getcare_province_id, 0)))

  JOIN (SELECT * FROM getcare_shipping_coupon_rule WHERE amount_gt<=@price_check AND (amount_lt IS NULL OR amount_lt>@price_check)) gscr ON getcare_shipping_coupon.id = gscr.getcare_shipping_coupon_id

WHERE getcare_shipping_coupon.getcare_vendor_id=@vendor_id AND getcare_shipping_coupon.active=1
  AND DATE(getcare_shipping_coupon.start_date)<=@date_check AND (getcare_shipping_coupon.end_date IS NULL OR DATE(getcare_shipping_coupon.end_date)>=@date_check)
  AND (gc.id IS NOT NULL OR (gscsc.id IS NULL AND gscc.id IS NULL));