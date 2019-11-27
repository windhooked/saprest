#https://answers.sap.com/questions/4484314/rhstrucget.html
HR='{\ 
   "exporting":{ 
      "plvar":"01",
      "otype":"O",
      "objid":"i_org_wa-org",
      "wegid":"A002",
      "begda":"i_org_wa-endda",
      "endda":"i_org_wa-endda",
      "pup_info":"X",
      "with_stext":"X",
      "tdepth":"0"
   },
   "tables":{ 
      "stru_tab":"h_stru_tab"
   },
   "exceptions":{ 
      "catalogue_problem":"1",
      "root_not_found":"2",
      "wegid_not_found":"3",
      "others":"4"
   }
}' 

test:
	curl -X POST -H "Content-Type: application/json" -d '{}' http://localhost:8088/rfc/STFC_STRUCTURE
hr:
	curl -X POST -H "Content-Type: application/json" -d $(HR) http://localhost:8088/rfc/RHPH_STRUCTURE_READ
