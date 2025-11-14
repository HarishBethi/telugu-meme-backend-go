package main

import (
	"encoding/json"
	"net/http"
)

type Template struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// CORS helper
func setHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func main() {

	// Root endpoint
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		setHeaders(w)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Go backend working!",
		})
	})

	// Templates endpoint
	http.HandleFunc("/templates", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodOptions {
			setHeaders(w)
			return
		}

		setHeaders(w)

		templates := []Template{

			// --- Your original templates ---
			{
				Name: "Brahmi_Krishna",
				URL:  "https://res.cloudinary.com/dui1jtybs/image/upload/v1763113242/brahmi-krishna_ptqsqu.gif?fl_cors",
			},
			{
				Name: "Arehoo1",
				URL:  "https://res.cloudinary.com/dui1jtybs/image/upload/v1763117753/arehoo1_w5azaj.jpg",
			},
			{
				Name: "Arehoo2",
				URL:  "https://res.cloudinary.com/dui1jtybs/image/upload/v1763117753/arehoo2_fhbxj4.jpg",
			},

			// --- All templates you provided ---
			{Name: "Samba_villian", URL: "https://res.cloudinary.com/dui1jtybs/image/upload/v1763117784/samba_villian_khnvrm.jpg"},
			{Name: "Venky_brahmi_shock3", URL: "https://res.cloudinary.com/dui1jtybs/image/upload/v1763117795/venky_brahmi_shock3_wwk50t.jpg"},
			{Name: "Venky_brahmi_shock2", URL: "https://res.cloudinary.com/dui1jtybs/image/upload/v1763117795/venky_brahmi_shock2_wurmbe.jpg"},
			{Name: "Venky_brahmi_shock", URL: "https://res.cloudinary.com/dui1jtybs/image/upload/v1763117795/venky_brahmi_shock_jokgax.jpg"},
			{Name: "Talent_chupisthunav_enty", URL: "https://res.cloudinary.com/dui1jtybs/image/upload/v1763117785/Talent_chupisthunav_enty_bl2esm.jpg"},
			{Name: "Tamanna_KIKK", URL: "https://res.cloudinary.com/dui1jtybs/image/upload/v1763117785/Tamanna_KIKK_sqrekc.jpg"},
			{Name: "Sontham_msn", URL: "https://res.cloudinary.com/dui1jtybs/image/upload/v1763117784/sontham_msn_qsx6jy.jpg"},
			{Name: "Veedu_musalodu_avvakudadhe", URL: "https://res.cloudinary.com/dui1jtybs/image/upload/v1763117785/veedu_musalodu_avvakudadhe_ugirbi.jpg"},
			{Name: "Sunil_andarivadu", URL: "https://res.cloudinary.com/dui1jtybs/image/upload/v1763117784/Sunil_andarivadu_uecet4.jpg"},
			{Name: "Venky_brahmi_innocent", URL: "https://res.cloudinary.com/dui1jtybs/image/upload/v1763117785/venky_brahmi_innocent_zpnmjr.jpg"},
			{Name: "Ready_brahmi", URL: "https://res.cloudinary.com/dui1jtybs/image/upload/v1763117784/ready_brahmi_jarpnj.jpg"},
			{Name: "Rajni_crying", URL: "https://res.cloudinary.com/dui1jtybs/image/upload/v1763117784/Rajni_crying_wmwqyy.jpg"},
			{Name: "Pruthvi_Atharinitiki_Daredi", URL: "https://res.cloudinary.com/dui1jtybs/image/upload/v1763117784/Pruthvi_Atharinitiki_Daredi_d2dhqd.jpg"},
			{Name: "Mari_antha_yadava", URL: "https://res.cloudinary.com/dui1jtybs/image/upload/v1763117775/mari_antha_yadava_sevzul.jpg"},
			{Name: "Neninthe_venumadhav", URL: "https://res.cloudinary.com/dui1jtybs/image/upload/v1763117775/neninthe_venumadhav_inrbmm.jpg"},
			{Name: "Manmadhudu_brahmi2", URL: "https://res.cloudinary.com/dui1jtybs/image/upload/v1763117774/Manmadhudu_brahmi2_ndqr2t.jpg"},
			{Name: "Natinchaku", URL: "https://res.cloudinary.com/dui1jtybs/image/upload/v1763117774/Natinchaku_igj0zg.jpg"},
			{Name: "Mr_bean", URL: "https://res.cloudinary.com/dui1jtybs/image/upload/v1763117774/mr_bean_mdaeqi.jpg"},
			{Name: "Manmadhudu_brahmi1", URL: "https://res.cloudinary.com/dui1jtybs/image/upload/v1763117774/Manmadhudu_brahmi1_j5lnqe.jpg"},
			{Name: "Krishna_brahmi", URL: "https://res.cloudinary.com/dui1jtybs/image/upload/v1763117774/krishna_brahmi_cqsyok.jpg"},
			{Name: "King_brahmi_6", URL: "https://res.cloudinary.com/dui1jtybs/image/upload/v1763117773/king_brahmi_6_qybabd.jpg"},
			{Name: "King_brahmi_7", URL: "https://res.cloudinary.com/dui1jtybs/image/upload/v1763117773/king_brahmi_7_vflsqw.jpg"},
			{Name: "King_brahmi_5", URL: "https://res.cloudinary.com/dui1jtybs/image/upload/v1763117762/king_brahmi_5_rxqkbq.jpg"},
			{Name: "King_brahmi_1", URL: "https://res.cloudinary.com/dui1jtybs/image/upload/v1763117755/king_brahmi_1_s1c6vv.jpg"},
			{Name: "King_brahmi_4", URL: "https://res.cloudinary.com/dui1jtybs/image/upload/v1763117755/king_brahmi_4_ksoand.jpg"},
			{Name: "King_brahmi_3", URL: "https://res.cloudinary.com/dui1jtybs/image/upload/v1763117755/king_brahmi_3_k1pqs6.jpg"},
			{Name: "Dubai_sreenu_krishna_baghavan", URL: "https://res.cloudinary.com/dui1jtybs/image/upload/v1763117754/dubai_sreenu_krishna_baghavan_pyuxzq.jpg"},
			{Name: "Dubai_sreenu_avs", URL: "https://res.cloudinary.com/dui1jtybs/image/upload/v1763117754/dubai_sreenu_avs_ozjazs.jpg"},
			{Name: "Brahmi_lakshmi_movie", URL: "https://res.cloudinary.com/dui1jtybs/image/upload/v1763117754/brahmi_lakshmi_movie_fypuo1.jpg"},
			{Name: "King_brahmi_2", URL: "https://res.cloudinary.com/dui1jtybs/image/upload/v1763117755/king_brahmi_2_fz6bwi.jpg"},
			{Name: "Brahmi_khaleja", URL: "https://res.cloudinary.com/dui1jtybs/image/upload/v1763117754/brahmi_khaleja_agip6g.jpg"},
			{Name: "Gamyam", URL: "https://res.cloudinary.com/dui1jtybs/image/upload/v1763117754/Gamyam_gvq44i.jpg"},
			{Name: "Emaindi_ra", URL: "https://res.cloudinary.com/dui1jtybs/image/upload/v1763117754/Emaindi_ra_dtbv9f.jpg"},
			{Name: "Brahmi_dhee", URL: "https://res.cloudinary.com/dui1jtybs/image/upload/v1763117754/brahmi_dhee_eywc8c.jpg"},
			{Name: "Brahmi_badhshah", URL: "https://res.cloudinary.com/dui1jtybs/image/upload/v1763117753/brahmi_badhshah_u0djlc.jpg"},
			{Name: "Brahmi_Attharitiniki_daredi", URL: "https://res.cloudinary.com/dui1jtybs/image/upload/v1763117753/Brahmi_Attharitiniki_daredi_w4ptlr.jpg"},
			{Name: "Bal_lav_de_ka_bal", URL: "https://res.cloudinary.com/dui1jtybs/image/upload/v1763117753/bal_lav_de_ka_bal_c2rco3.jpg"},
			{Name: "Apartichitudu_prakashraj", URL: "https://res.cloudinary.com/dui1jtybs/image/upload/v1763117753/apartichitudu_prakashraj_j7hzs4.jpg"},
			{Name: "Ali_Super_movie", URL: "https://res.cloudinary.com/dui1jtybs/image/upload/v1763117753/Ali_Super_movie_qsa758.jpg"},
			{Name: "Addi_movie_villian", URL: "https://res.cloudinary.com/dui1jtybs/image/upload/v1763117753/Addi_movie_villian_bqqnaj.jpg"},
		}

		json.NewEncoder(w).Encode(templates)
	})

	http.ListenAndServe(":8080", nil)
}
