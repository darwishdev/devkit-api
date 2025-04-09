package adapter

import (
	"encoding/json"
	"time"

	"github.com/darwishdev/devkit-api/db"
	devkitv1 "github.com/darwishdev/devkit-api/proto_gen/devkit/v1"
	"github.com/rs/zerolog/log"
)

func (a *TenantAdapter) TenantEntityGrpcFromSql(resp *db.TenantsSchemaTenant) *devkitv1.TenantsSchemaTenant {
	var tenantLinks map[string]string
	log.Debug().Interface("tenant links is", resp.TenantLinks).Msg("links is")
	if len(resp.TenantLinks) > 0 {
		if err := json.Unmarshal(resp.TenantLinks, &tenantLinks); err == nil {
			log.Error().Err(err).Msg("error parsing partial links")
		}
	}

	return &devkitv1.TenantsSchemaTenant{
		TenantId:               int32(resp.TenantID),
		TenantName:             resp.TenantName,
		TenantNameAr:           resp.TenantNameAr.String,
		TenantPhone:            resp.TenantPhone.String,
		TenantAddress:          resp.TenantAddress.String,
		TenantAddressAr:        resp.TenantAddressAr.String,
		TenantDescription:      resp.TenantDescription.String,
		TenantDescriptionAr:    resp.TenantDescriptionAr.String,
		TenantEmail:            resp.TenantEmail.String,
		TenantLinks:            tenantLinks,
		TenantLogo:             resp.TenantLogo.String,
		TenantValues:           resp.TenantValues.String,
		TenantVision:           resp.TenantVision.String,
		TenantMission:          resp.TenantMission.String,
		TenantLogoVertical:     resp.TenantLogoVertical.String,
		TenantLogoDark:         resp.TenantLogoDark.String,
		TenantLogoDarkVertical: resp.TenantLogoDarkVertical.String,
		CreatedAt:              db.TimeToString(resp.CreatedAt.Time),
		DeletedAt:              db.TimeToString(resp.DeletedAt.Time),
	}
}

func (a *TenantAdapter) TenantCreateUpdateSqlFromGrpc(req *devkitv1.TenantCreateUpdateRequest) *db.TenantCreateUpdateParams {
	tenantLinks, err := json.Marshal(req.GetTenantLinks())
	if err != nil {
		log.Error().Err(err).Msg("error parsing partial links")
	}

	resp := &db.TenantCreateUpdateParams{
		TenantID:               req.TenantId,
		TenantName:             req.TenantName,
		TenantNameAr:           req.TenantNameAr,
		TenantLinks:            tenantLinks,
		TenantPhone:            req.TenantPhone,
		TenantAddress:          req.TenantAddress,
		TenantAddressAr:        req.TenantAddressAr,
		TenantValues:           req.TenantValues,
		TenantVision:           req.TenantVision,
		TenantMission:          req.TenantMission,
		TenantEmail:            req.TenantEmail,
		TenantDescription:      req.TenantDescription,
		TenantDescriptionAr:    req.TenantDescriptionAr,
		TenantLogo:             req.TenantLogo,
		TenantLogoVertical:     req.TenantLogoVertical,
		TenantLogoDark:         req.TenantLogoDark,
		TenantLogoDarkVertical: req.TenantLogoDarkVertical,
	}
	return resp
}
func (a *TenantAdapter) TenantListGrpcFromSql(resp *[]db.TenantsSchemaTenant) *devkitv1.TenantListResponse {
	records := make([]*devkitv1.TenantsSchemaTenant, 0)
	deletedRecords := make([]*devkitv1.TenantsSchemaTenant, 0)
	for _, v := range *resp {
		record := a.TenantEntityGrpcFromSql(&v)
		if v.DeletedAt.Valid {
			deletedRecords = append(deletedRecords, record)
		} else {
			records = append(records, record)
		}
	}
	return &devkitv1.TenantListResponse{
		DeletedRecords: deletedRecords,
		Records:        records,
	}
}
func (a *TenantAdapter) TenantFindGrpcFromSql(resp *db.TenantFindRow) *devkitv1.TenantFindResponse {

	var navigations []*devkitv1.NavigationBar
	var tenantLinks map[string]string
	if len(resp.TenantLinks) > 0 {
		if err := json.Unmarshal(resp.TenantLinks, &tenantLinks); err == nil {
			log.Error().Err(err).Msg("error parsing partial links")
		}
	}

	log.Debug().Interface("debug", tenantLinks).Msg("navigations are")

	if len(resp.Navigations) > 0 {
		err := json.Unmarshal(resp.Navigations, &navigations)
		log.Error().Err(err).Msg("error parsing the navigations")
	}
	var pages []*devkitv1.TenantsSchemaPageView
	if len(resp.Pages) > 0 {
		err := json.Unmarshal(resp.Pages, &pages)
		log.Error().Err(err).Msg("error parsing the pages")
	}
	return &devkitv1.TenantFindResponse{
		Tenant: &devkitv1.TenantsSchemaTenantView{
			TenantId:               resp.TenantID,
			TenantName:             resp.TenantName,
			TenantNameAr:           resp.TenantNameAr.String,
			TenantPhone:            resp.TenantPhone.String,
			TenantAddress:          resp.TenantAddress.String,
			TenantAddressAr:        resp.TenantAddressAr.String,
			TenantDescription:      resp.TenantDescription.String,
			TenantDescriptionAr:    resp.TenantDescriptionAr.String,
			TenantLinks:            tenantLinks,
			TenantEmail:            resp.TenantEmail.String,
			TenantLogo:             resp.TenantLogo.String,
			TenantLogoVertical:     resp.TenantLogoVertical.String,
			TenantLogoDark:         resp.TenantLogoDark.String,
			TenantLogoDarkVertical: resp.TenantLogoDarkVertical.String,
			TenantValues:           resp.TenantValues.String,
			TenantVision:           resp.TenantVision.String,
			TenantMission:          resp.TenantMission.String,
			CreatedAt:              resp.CreatedAt.Time.Format(time.RFC3339),
			UpdatedAt:              resp.UpdatedAt.Time.Format(time.RFC3339),
			DeletedAt:              resp.DeletedAt.Time.Format(time.RFC3339),
			Pages:                  pages,
			Navigations:            navigations,
		},
	}
}
func (a *TenantAdapter) TenantDeleteRestoreGrpcFromSql(resp *[]db.TenantsSchemaTenant) *devkitv1.TenantDeleteRestoreResponse {
	records := make([]*devkitv1.TenantsSchemaTenant, 0)
	for _, v := range *resp {
		record := a.TenantEntityGrpcFromSql(&v)
		records = append(records, record)
	}
	return &devkitv1.TenantDeleteRestoreResponse{
		Records: records,
	}

}
