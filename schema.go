package main

type BambuLabsH2S struct {
	Print struct {
		TwoD struct {
			Bs struct {
				Bi []struct {
					ClockIn   bool          `json:"clock_in"`
					EstTime   int           `json:"est_time"`
					Idx       int           `json:"idx"`
					PrintThen bool          `json:"print_then"`
					ProcList  []interface{} `json:"proc_list"`
					StepType  int           `json:"step_type"`
					ToolInfo  struct {
						Color    string  `json:"color"`
						Diameter float64 `json:"diameter"`
						ID       int     `json:"id"`
					} `json:"tool_info"`
					Type int `json:"type"`
				} `json:"bi"`
				TotalTime int `json:"total_time"`
			} `json:"bs"`
			Cond     int `json:"cond"`
			CurStage struct {
				ClockInTime int `json:"clock_in_time"`
				Idx         int `json:"idx"`
				LeftTime    int `json:"left_time"`
				Process     int `json:"process"`
				State       int `json:"state"`
			} `json:"cur_stage"`
			FirstConfirm bool `json:"first_confirm"`
			Makeable     bool `json:"makeable"`
			Material     struct {
				CurIDList []interface{} `json:"cur_id_list"`
				State     int           `json:"state"`
				TarID     string        `json:"tar_id"`
				TarName   string        `json:"tar_name"`
			} `json:"material"`
		} `json:"2D"`
		ThreeD struct {
			LayerNum      int `json:"layer_num"`
			TotalLayerNum int `json:"total_layer_num"`
		} `json:"3D"`
		Ams struct {
			Ams []struct {
				DryTime     int    `json:"dry_time"`
				Humidity    string `json:"humidity"`
				HumidityRaw string `json:"humidity_raw"`
				ID          string `json:"id"`
				Info        string `json:"info"`
				Temp        string `json:"temp"`
				Tray        []struct {
					BedTemp       string   `json:"bed_temp,omitempty"`
					BedTempType   string   `json:"bed_temp_type,omitempty"`
					CaliIdx       int      `json:"cali_idx,omitempty"`
					Cols          []string `json:"cols,omitempty"`
					Ctype         int      `json:"ctype,omitempty"`
					DryingTemp    string   `json:"drying_temp,omitempty"`
					DryingTime    string   `json:"drying_time,omitempty"`
					ID            string   `json:"id"`
					NozzleTempMax string   `json:"nozzle_temp_max,omitempty"`
					NozzleTempMin string   `json:"nozzle_temp_min,omitempty"`
					Remain        int      `json:"remain,omitempty"`
					State         int      `json:"state"`
					TagUID        string   `json:"tag_uid,omitempty"`
					TotalLen      int      `json:"total_len,omitempty"`
					TrayColor     string   `json:"tray_color,omitempty"`
					TrayDiameter  string   `json:"tray_diameter,omitempty"`
					TrayIDName    string   `json:"tray_id_name,omitempty"`
					TrayInfoIdx   string   `json:"tray_info_idx,omitempty"`
					TraySubBrands string   `json:"tray_sub_brands,omitempty"`
					TrayType      string   `json:"tray_type,omitempty"`
					TrayUUID      string   `json:"tray_uuid,omitempty"`
					TrayWeight    string   `json:"tray_weight,omitempty"`
					XcamInfo      string   `json:"xcam_info,omitempty"`
				} `json:"tray"`
			} `json:"ams"`
			AmsExistBits     string `json:"ams_exist_bits"`
			AmsExistBitsRaw  string `json:"ams_exist_bits_raw"`
			CaliID           int    `json:"cali_id"`
			CaliStat         int    `json:"cali_stat"`
			InsertFlag       bool   `json:"insert_flag"`
			PowerOnFlag      bool   `json:"power_on_flag"`
			TrayExistBits    string `json:"tray_exist_bits"`
			TrayIsBblBits    string `json:"tray_is_bbl_bits"`
			TrayNow          string `json:"tray_now"`
			TrayPre          string `json:"tray_pre"`
			TrayReadDoneBits string `json:"tray_read_done_bits"`
			TrayReadingBits  string `json:"tray_reading_bits"`
			TrayTar          string `json:"tray_tar"`
			UnbindAmsStat    int    `json:"unbind_ams_stat"`
			Version          int    `json:"version"`
		} `json:"ams"`
		AmsRfidStatus   int     `json:"ams_rfid_status"`
		AmsStatus       int     `json:"ams_status"`
		ApErr           int     `json:"ap_err"`
		Aux             string  `json:"aux"`
		AuxPartFan      bool    `json:"aux_part_fan"`
		BatchID         int     `json:"batch_id"`
		BedTargetTemper float64 `json:"bed_target_temper"`
		BedTemper       float64 `json:"bed_temper"`
		BigFan1Speed    string  `json:"big_fan1_speed"`
		BigFan2Speed    string  `json:"big_fan2_speed"`
		CaliVersion     int     `json:"cali_version"`
		CanvasID        int     `json:"canvas_id"`
		Care            []struct {
			ID   string `json:"id"`
			Info string `json:"info"`
		} `json:"care"`
		Cfg             string `json:"cfg"`
		Command         string `json:"command"`
		CoolingFanSpeed string `json:"cooling_fan_speed"`
		DesignID        string `json:"design_id"`
		Device          struct {
			Airduct struct {
				ModeCur  int `json:"modeCur"`
				ModeList []struct {
					Ctrl   []int `json:"ctrl"`
					ModeID int   `json:"modeId"`
					Off    []int `json:"off"`
				} `json:"modeList"`
				Parts []struct {
					Func  int `json:"func"`
					ID    int `json:"id"`
					Range int `json:"range"`
					State int `json:"state"`
				} `json:"parts"`
				SubMode int `json:"subMode"`
			} `json:"airduct"`
			Bed struct {
				Info struct {
					Temp int `json:"temp"`
				} `json:"info"`
				State int `json:"state"`
			} `json:"bed"`
			BedTemp int `json:"bed_temp"`
			Cam     struct {
				Laser struct {
					Cond  int `json:"cond"`
					State int `json:"state"`
				} `json:"laser"`
			} `json:"cam"`
			Ctc struct {
				Info struct {
					Temp int `json:"temp"`
				} `json:"info"`
				State int `json:"state"`
			} `json:"ctc"`
			ExtTool struct {
				Calib   int    `json:"calib"`
				LowPrec bool   `json:"low_prec"`
				Mount   int    `json:"mount"`
				Mount3D int    `json:"mount_3d"`
				ThTemp  int    `json:"th_temp"`
				Type    string `json:"type"`
			} `json:"ext_tool"`
			Extruder struct {
				Info []struct {
					FilamBak []interface{} `json:"filam_bak"`
					Hnow     int           `json:"hnow"`
					Hpre     int           `json:"hpre"`
					Htar     int           `json:"htar"`
					ID       int           `json:"id"`
					Info     int           `json:"info"`
					Snow     int           `json:"snow"`
					Spre     int           `json:"spre"`
					Star     int           `json:"star"`
					Stat     int           `json:"stat"`
					Temp     int           `json:"temp"`
				} `json:"info"`
				State int `json:"state"`
			} `json:"extruder"`
			Fan     int `json:"fan"`
			FireExt struct {
				Cd    int `json:"cd"`
				State int `json:"state"`
			} `json:"fire_ext"`
			Laser struct {
				Power int `json:"power"`
			} `json:"laser"`
			Nozzle struct {
				Exist int `json:"exist"`
				Info  []struct {
					Diameter float64 `json:"diameter"`
					ID       int     `json:"id"`
					Tm       int     `json:"tm"`
					Type     string  `json:"type"`
					Wear     int     `json:"wear"`
				} `json:"info"`
				State int `json:"state"`
			} `json:"nozzle"`
			Plate struct {
				Base     int    `json:"base"`
				Cali2DID string `json:"cali2d_id"`
				CurID    string `json:"cur_id"`
				Mat      int    `json:"mat"`
				TarID    string `json:"tar_id"`
			} `json:"plate"`
			Type int `json:"type"`
		} `json:"device"`
		Err                     string `json:"err"`
		FailReason              string `json:"fail_reason"`
		FanGear                 int    `json:"fan_gear"`
		File                    string `json:"file"`
		ForceUpgrade            bool   `json:"force_upgrade"`
		Fun                     string `json:"fun"`
		GcodeFile               string `json:"gcode_file"`
		GcodeFilePreparePercent string `json:"gcode_file_prepare_percent"`
		GcodeState              string `json:"gcode_state"`
		HeatbreakFanSpeed       string `json:"heatbreak_fan_speed"`
		Hms                     []struct {
			Attr int `json:"attr"`
			Code int `json:"code"`
		} `json:"hms"`
		HomeFlag      int `json:"home_flag"`
		HwSwitchState int `json:"hw_switch_state"`
		Info          struct {
			Temp int `json:"temp"`
		} `json:"info"`
		Ipcam struct {
			AgoraService    string `json:"agora_service"`
			BrtcService     string `json:"brtc_service"`
			BsState         int    `json:"bs_state"`
			IpcamDev        string `json:"ipcam_dev"`
			IpcamRecord     string `json:"ipcam_record"`
			LaserPreviewRes int    `json:"laser_preview_res"`
			ModeBits        int    `json:"mode_bits"`
			Resolution      string `json:"resolution"`
			RtspURL         string `json:"rtsp_url"`
			Timelapse       string `json:"timelapse"`
			TlStoreHpdType  int    `json:"tl_store_hpd_type"`
			TlStorePathType int    `json:"tl_store_path_type"`
			TutkServer      string `json:"tutk_server"`
		} `json:"ipcam"`
		Job struct {
			CurStage struct {
				Idx   int `json:"idx"`
				State int `json:"state"`
			} `json:"cur_stage"`
			Stage []struct {
				ClockIn   bool          `json:"clock_in"`
				Color     []string      `json:"color"`
				Diameter  []float64     `json:"diameter"`
				EstTime   int           `json:"est_time"`
				Heigh     float64       `json:"heigh"`
				Idx       int           `json:"idx"`
				Platform  string        `json:"platform"`
				PrintThen bool          `json:"print_then"`
				ProcList  []interface{} `json:"proc_list"`
				Tool      []string      `json:"tool"`
				Type      int           `json:"type"`
			} `json:"stage"`
		} `json:"job"`
		JobAttr      int    `json:"job_attr"`
		JobID        string `json:"job_id"`
		LanTaskID    string `json:"lan_task_id"`
		LayerNum     int    `json:"layer_num"`
		LightsReport []struct {
			Mode string `json:"mode"`
			Node string `json:"node"`
		} `json:"lights_report"`
		Mapping          []int  `json:"mapping"`
		McAction         int    `json:"mc_action"`
		McErr            int    `json:"mc_err"`
		McPercent        int    `json:"mc_percent"`
		McPrintErrorCode string `json:"mc_print_error_code"`
		McPrintStage     string `json:"mc_print_stage"`
		McPrintSubStage  int    `json:"mc_print_sub_stage"`
		McRemainingTime  int    `json:"mc_remaining_time"`
		McStage          int    `json:"mc_stage"`
		ModelID          string `json:"model_id"`
		Net              struct {
			Conf int `json:"conf"`
			Info []struct {
				IP   int64 `json:"ip"`
				Mask int   `json:"mask"`
			} `json:"info"`
		} `json:"net"`
		NozzleDiameter     string  `json:"nozzle_diameter"`
		NozzleTargetTemper float64 `json:"nozzle_target_temper"`
		NozzleTemper       float64 `json:"nozzle_temper"`
		NozzleType         string  `json:"nozzle_type"`
		Online             struct {
			Ahb     bool `json:"ahb"`
			Version int  `json:"version"`
		} `json:"online"`
		Percent          int           `json:"percent"`
		PlateCnt         int           `json:"plate_cnt"`
		PlateID          int           `json:"plate_id"`
		PlateIdx         int           `json:"plate_idx"`
		PreparePer       int           `json:"prepare_per"`
		PrintError       int           `json:"print_error"`
		PrintGcodeAction int           `json:"print_gcode_action"`
		PrintRealAction  int           `json:"print_real_action"`
		PrintType        string        `json:"print_type"`
		ProfileID        string        `json:"profile_id"`
		ProjectID        string        `json:"project_id"`
		Queue            int           `json:"queue"`
		QueueEst         int           `json:"queue_est"`
		QueueNumber      int           `json:"queue_number"`
		QueueSts         int           `json:"queue_sts"`
		QueueTotal       int           `json:"queue_total"`
		RemainTime       int           `json:"remain_time"`
		SObj             []interface{} `json:"s_obj"`
		Sdcard           bool          `json:"sdcard"`
		SequenceID       string        `json:"sequence_id"`
		SpdLvl           int           `json:"spd_lvl"`
		SpdMag           int           `json:"spd_mag"`
		Stat             string        `json:"stat"`
		State            int           `json:"state"`
		Stg              []int         `json:"stg"`
		StgCur           int           `json:"stg_cur"`
		SubtaskID        string        `json:"subtask_id"`
		SubtaskName      string        `json:"subtask_name"`
		TaskID           string        `json:"task_id"`
		TotalLayerNum    int           `json:"total_layer_num"`
		UpgradeState     struct {
			AhbNewVersionNumber string `json:"ahb_new_version_number"`
			AmsNewVersionNumber string `json:"ams_new_version_number"`
			ConsistencyRequest  bool   `json:"consistency_request"`
			DisState            int    `json:"dis_state"`
			ErrCode             int    `json:"err_code"`
			ExtNewVersionNumber string `json:"ext_new_version_number"`
			ForceUpgrade        bool   `json:"force_upgrade"`
			Idx                 int    `json:"idx"`
			Idx2                int    `json:"idx2"`
			LowerLimit          string `json:"lower_limit"`
			Message             string `json:"message"`
			Module              string `json:"module"`
			NewVersionState     int    `json:"new_version_state"`
			OtaNewVersionNumber string `json:"ota_new_version_number"`
			Progress            string `json:"progress"`
			SequenceID          int    `json:"sequence_id"`
			Sn                  string `json:"sn"`
			Status              string `json:"status"`
		} `json:"upgrade_state"`
		Upload struct {
			FileSize      int    `json:"file_size"`
			FinishSize    int    `json:"finish_size"`
			Message       string `json:"message"`
			OssURL        string `json:"oss_url"`
			Progress      int    `json:"progress"`
			SequenceID    string `json:"sequence_id"`
			Speed         int    `json:"speed"`
			Status        string `json:"status"`
			TaskID        string `json:"task_id"`
			TimeRemaining int    `json:"time_remaining"`
			TroubleID     string `json:"trouble_id"`
		} `json:"upload"`
		Ver     string `json:"ver"`
		VirSlot []struct {
			BedTemp       string   `json:"bed_temp"`
			BedTempType   string   `json:"bed_temp_type"`
			CaliIdx       int      `json:"cali_idx"`
			Cols          []string `json:"cols"`
			Ctype         int      `json:"ctype"`
			DryingTemp    string   `json:"drying_temp"`
			DryingTime    string   `json:"drying_time"`
			ID            string   `json:"id"`
			NozzleTempMax string   `json:"nozzle_temp_max"`
			NozzleTempMin string   `json:"nozzle_temp_min"`
			Remain        int      `json:"remain"`
			TagUID        string   `json:"tag_uid"`
			TotalLen      int      `json:"total_len"`
			TrayColor     string   `json:"tray_color"`
			TrayDiameter  string   `json:"tray_diameter"`
			TrayIDName    string   `json:"tray_id_name"`
			TrayInfoIdx   string   `json:"tray_info_idx"`
			TraySubBrands string   `json:"tray_sub_brands"`
			TrayType      string   `json:"tray_type"`
			TrayUUID      string   `json:"tray_uuid"`
			TrayWeight    string   `json:"tray_weight"`
			XcamInfo      string   `json:"xcam_info"`
		} `json:"vir_slot"`
		WifiSignal string `json:"wifi_signal"`
		Xcam       struct {
			AllowSkipParts           bool   `json:"allow_skip_parts"`
			BuildplateMarkerDetector bool   `json:"buildplate_marker_detector"`
			Cfg                      int    `json:"cfg"`
			FirstLayerInspector      bool   `json:"first_layer_inspector"`
			HaltPrintSensitivity     string `json:"halt_print_sensitivity"`
			PrintHalt                bool   `json:"print_halt"`
			PrintingMonitor          bool   `json:"printing_monitor"`
			SpaghettiDetector        bool   `json:"spaghetti_detector"`
		} `json:"xcam"`
		XcamStatus string `json:"xcam_status"`
	} `json:"print"`
}
