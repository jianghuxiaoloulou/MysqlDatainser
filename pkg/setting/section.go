package setting

type GeneralSettingS struct {
	LogSavePath string
	LogFileName string
	LogFileExt  string
	LogMaxSize  int
	LogMaxAge   int
}

type DatabaseSettingS struct {
	DBConn       string
	DBType       string
	MaxIdleConns int
	MaxOpenConns int
}

func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}
	return nil
}
