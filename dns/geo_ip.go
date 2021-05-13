package dns

import (
	"github.com/demonoid81/proxyM/assets"
	"github.com/demonoid81/proxyM/log"
	"github.com/oschwald/geoip2-golang"
	"net"
)

var geoipDB *geoip2.Reader

func InitGeoIP(dbFile string) error {
	var err error
	geoipFileBytes, err := assets.ReadFile(dbFile)
	if err != nil {
		log.Logger.Errorf("[GeoIP] read failed [%v]", err)
		return err
	}
	geoipDB, err = geoip2.FromBytes(geoipFileBytes)
	if err != nil {
		return err
	}
	return nil
}

func GeoLookUp(ip string) (countryCode string) {
	if geoipDB == nil {
		return
	}
	netIP := net.ParseIP(ip)
	if netIP == nil {
		return
	}
	country, err := geoipDB.Country(netIP)
	if err == nil && country != nil {
		log.Logger.Debugf("[GeoIP] lookup [%s] country -> [%s]", ip, country.Country.IsoCode)
		return country.Country.IsoCode
	}
	log.Logger.Debugf("[GeoIP] lookup [%s] country failed: %s", ip, err.Error())
	return
}

func CloseGeoDB() error {
	return geoipDB.Close()
}
