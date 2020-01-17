#CALL FUNCTION 'BAPI_MATERIAL_AVAILABILITY' "ATP information
#  EXPORTING
#    plant =                     " bapimatvp-werks  Plant
#    material =                  " bapimatvp-matnr  Material number
#    unit =                      " bapiadmm-unit  Unit of measure for display
#*   check_rule =                " bapit441v-prreg  Checking rule
#*   stge_loc =                  " bapicm61v-lgort  Storage location
#*   batch =                     " bapicm61v-charg  Batch
#*   customer =                  " bapiknvvky-customer  Customer number
#*   doc_number =                " bapisdhead-doc_number  Document number
#*   itm_number =                " bapisditm-itm_number  Item number
#*   wbs_elem =                  " bapisditm-wbs_elem  WBS Element
#*   stock_ind =                 " bapicm61v-sobkz  Special Stock Indicator
#*   dec_for_rounding =          " bapicm61m-andec  No. of decimal places to which rounding should be performed
#*   dec_for_rounding_x =        " bapiupdate    Updated information in related user data field
#*   read_atp_lock =             " bapicm61v-vfpst  Control indicator for availability check
#*   read_atp_lock_x =           " bapiupdate    Updated information in related user data field
#*   material_evg =              " bapimgvmatnr
#  IMPORTING
#    endleadtme =                " bapicm61m-wzter  End of replenishment lead time
#    av_qty_plt =                " bapicm61v-wkbst  Quantity available at plant level
#    dialogflag =                " bapicm61v-diafl  Indicator (X = not available, N = no check)
#    return =                    " bapireturn
#  TABLES
#    wmdvsx =                    " bapiwmdvs     Input table (date and quantity)
#    wmdvex =                    " bapiwmdve     Output table (date and ATP quantity)
#    .  "  BAPI_MATERIAL_AVAILABILITY
curl -X POST -H "Content-Type: application/json" -d $(MM) http://localhost:8088/rfc/BAPI_MATERIAL_AVAILABILITY

