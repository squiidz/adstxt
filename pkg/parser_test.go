package pkg

import (
	"strings"
	"testing"
)

func TestParseAdsTxt(t *testing.T) {
	adstxt := strings.NewReader(`
		#this is a comment
		taboola.com,1201816,DIRECT,c228e6794e811952
		taboola.com,1201817,DIRECT # test comment 
		taboola.com,1201818 , DIRECT ,
		taboola.com,1201819,DIRECT # comment with comma, 1,2,3
		<invalid></invalid>
	`)
	pub, err := parseAdsTxt("test.com", adstxt)
	if err != nil {
		t.Error(err)
	}
	if len(pub.Sellers) != 4 {
		t.Fail()
	}
}

func TestParseAdsTxtWithComment(t *testing.T) {
	adstxt := strings.NewReader(`
		taboola.com,1201819,DIRECT # comment with comma, 1,2,3
	`)
	pub, err := parseAdsTxt("test.com", adstxt)
	if err != nil {
		t.Error(err)
	}
	if pub.Sellers[0].TypeOfAccount != "DIRECT" {
		t.Fail()
	}
	if pub.Sellers[0].CertAuthID != "" {
		t.Fail()
	}
}

func TestParseAdsTxtWithComments(t *testing.T) {
	adstxt := strings.NewReader(`
		# Comment
		taboola.com,1201819,DIRECT # comment with comma, 1,2,3 # 2
		# Comment,
	`)
	pub, err := parseAdsTxt("test.com", adstxt)
	if err != nil {
		t.Error(err)
	}
	if pub.Sellers[0].TypeOfAccount != "DIRECT" {
		t.Fail()
	}
	if pub.Sellers[0].CertAuthID != "" {
		t.Fail()
	}
}

func TestParseAds2Txt(t *testing.T) {
	adstxt := strings.NewReader(`
		#this is a comment
		taboola.com,1201816,DIRECT,c228e6794e811952
		taboola.com,1201817,DIRECT # test comment 
		taboola.com,1201818 , DIRECT ,
		taboola.com,1201819,DIRECT # comment with comma, 1,2,3
		<invalid></invalid>
	`)
	pub, err := parseAdsTxt2("test.com", adstxt)
	if err != nil {
		t.Error(err)
	}
	if len(pub.Sellers) != 4 {
		t.Fail()
	}
}

func TestParseAds2TxtWithComment(t *testing.T) {
	adstxt := strings.NewReader(`
		taboola.com,1201819,DIRECT # comment with comma, 1,2,3
	`)
	pub, err := parseAdsTxt2("test.com", adstxt)
	if err != nil {
		t.Error(err)
	}
	if pub.Sellers[0].TypeOfAccount != "DIRECT" {
		t.Fail()
	}
	if pub.Sellers[0].CertAuthID != "" {
		t.Fail()
	}
}

func TestParseAds2TxtWithComments(t *testing.T) {
	adstxt := strings.NewReader(`
		# Comment
		taboola.com,1201819,DIRECT # comment with comma, 1,2,3 # 2
		# Comment,
	`)
	pub, err := parseAdsTxt2("test.com", adstxt)
	if err != nil {
		t.Error(err)
	}
	if pub.Sellers[0].TypeOfAccount != "DIRECT" {
		t.Fail()
	}
	if pub.Sellers[0].CertAuthID != "" {
		t.Fail()
	}
}

func BenchmarkParseAds(b *testing.B) {
	for i := 0; i < b.N; i++ {
		parseAdsTxt("test", adstxtblob)
	}
}

func BenchmarkParseAds2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		parseAdsTxt2("test", adstxtblob)
	}
}

var adstxtblob = strings.NewReader(`
#FMG

amazon-adsystem.com, 3076, DIRECT
facebook.com, 743604429120938, DIRECT
rubiconproject.com, 12156, DIRECT, 0bfd66d529a55807
rubiconproject.com, 209082, DIRECT, 0bfd66d529a55807
rubiconproject.com, 17702, DIRECT, 0bfd66d529a55807
indexexchange.com, 183957, DIRECT
indexexchange.com, 184856, DIRECT
adtech.com, 10434, DIRECT
aolcloud.net, 10434, DIRECT
advertising.com, 10809, DIRECT
google.com, pub-9268440883448925, DIRECT
yieldmo.com, Fusion%20Media%20Group, DIRECT
yieldmo.com, 1701426062972061316, DIRECT
openx.com, 539169203, DIRECT
kargo.com, 200, DIRECT
indexexchange.com, 184081, RESELLER
triplelift.com, 2480, DIRECT, 6c33edb13117fd86
Criteo.com, 129974, DIRECT
spotxchange.com, 229777, DIRECT, 7842df1d2fe2db34
spotx.tv, 229777, DIRECT, 7842df1d2fe2db34
rubiconproject.com, 17740, DIRECT, 0bfd66d529a55807
yieldmo.com, 2150098008502076779, DIRECT
appnexus.com, 7656, DIRECT
taboola.com,1201816,DIRECT,c228e6794e811952
taboola.com,1201804,DIRECT,c228e6794e811952
taboola.com,1201800,DIRECT,c228e6794e811952
taboola.com,1201822,DIRECT,c228e6794e811952
taboola.com,1201759,DIRECT,c228e6794e811952
taboola.com,1201806,DIRECT,c228e6794e811952
taboola.com,1201801,DIRECT,c228e6794e811952
taboola.com,1201798,DIRECT,c228e6794e811952
taboola.com,1201766,DIRECT,c228e6794e811952
taboola.com,1201819,DIRECT,c228e6794e811952
taboola.com,1201762,DIRECT,c228e6794e811952
taboola.com,1201797,DIRECT,c228e6794e811952
taboola.com,1201799,DIRECT,c228e6794e811952
taboola.com,1201757,DIRECT,c228e6794e811952
taboola.com,1201823,DIRECT,c228e6794e811952
taboola.com,1201803,DIRECT,c228e6794e811952
spotx.tv,71451,RESELLER
spotxchange.com, 71451, RESELLER
advertising.com, 8603, RESELLER
pubmatic.com, 156307, RESELLER, 5d62403b186f2ace
appnexus.com, 3364, RESELLER
Indexexchange.com, 183756, RESELLER
contextweb.com, 560382, RESELLER
openx.com, 539154393, RESELLER
tremorhub.com, z87wm, RESELLER, 1a4e959a1b50034a
rubiconproject.com, 16698, RESELLER, 0bfd66d529a55807
freewheel.tv, 799841, RESELLER 
freewheel.tv, 799921, RESELLER 
aol.com, 53392, RESELLER 
rhythmone.com, 1166984029, RESELLER, a670c89d4a324e47
teads.tv, 18530, DIRECT, 15a9c44f6d26cbe1
behave.com, 1062, Direct
behave.com, 1063, Direct
behave.com, 1064, Direct
behave.com, 1065, Direct
behave.com, 1066, Direct
indexexchange.com, 183753, RESELLER
pubmatic.com, 156512, RESELLER
media.net, 8CU74RYRS, DIRECT
adtech.com, 11716, DIRECT, e1a5b5b6e3255540 
indexexchange.com, 175407, RESELLER, 50b1c356f2c5c8fc
openx.com, 537143344, RESELLER
pubmatic.com, 156078, RESELLER, 5d62403b186f2ace
contextweb.com, 558299, RESELLER, 89ff185a4c4e857c
openx.com, 540842041, DIRECT, 6a698e2ec38604c6
synacor.com, 82212, DIRECT, e108f11b2cdf7d5b
advertising.com, 19623, RESELLER # AOL - One
appnexus.com, 310, RESELLER, f5ab79cb980f11d1 # AppNexus
appnexus.com, 9316, RESELLER, f5ab79cb980f11d1 # AppNexus
pubmatic.com, 156344, RESELLER, 5d62403b186f2ace # Pubmatic
rubiconproject.com, 13344, RESELLER, 0bfd66d529a55807 # Rubicon
springserve.com, 278, RESELLER, a24eb641fc82e93d # SpringServe
kargo.com, 200 , direct
indexexchange.com, 184081 , reseller
contextweb.com, 562001, RESELLER, 89ff185a4c4e857c
appnexus.com, 8173, RESELLER
consumable.com, 2000904, DIRECT
adtech.com, 10947, DIRECT
appnexus.com, 7556, DIRECT
google.com, pub-6694481294649483, DIRECT
indexexchange.com, 184914, DIRECT
indexexchange.com, 186248, DIRECT, 50b1c356f2c5c8fc
indexexchange.com, 187454, DIRECT
openx.com, 537150004, DIRECT
openx.com, 539699341, DIRECT, 6a698e2ec38604c6
pubmatic.com, 156319, DIRECT
rubiconproject.com, 17632, DIRECT, 0bfd66d529a55807
rubiconproject.com, 18890, DIRECT, 0bfd66d529a55807
trustx.org, 76, DIRECT

#International

pubmatic.com, 156292, Direct
google.com, pub-3560785102412960, DIRECT, f08c47fec0942fa0
google.com, pub-1417235232628100, DIRECT, f08c47fec0942fa0
rubiconproject.com, 9788, DIRECT, 0bfd66d529a55807
openx.com, 537151514, DIRECT, a698e2ec38604c6
openx.com, 539136001, DIRECT, a698e2ec38604c6
appnexus.com, 7402, DIRECT
rhythmone.com,1126204936,DIRECT,a670c89d4a324e47
rhythmone.com,1804546489,DIRECT,a670c89d4a324e49
rhythmone.com,1164793360,DIRECT,a670c89d4a324e50
rhythmone.com,3273850690,DIRECT,a670c89d4a324e53
indexexchange.com, 185306, DIRECT
adyoulike.com, 594e25c150d724d1ec4d188b323da403, DIRECT
sharethrough.com, 9ef5c2ea, DIRECT, d53b998a7bd4ecd2
undertone.com, 2187, DIRECT
appnexus.com, 2234, RESELLER
openx.com, 537153564, RESELLER
loopme.com, 6444, direct, 6c8d5f95897a5a3b
sharethrough.com, 9ef5c2ea, DIRECT, d53b998a7bd4ecd2
Appnexus.com, 564, RESELLER, 912771
google.com, pub-1599446853501333, DIRECT, f08c47fec0942fa0
taboola.com,1201816,DIRECT,c228e6794e811952
taboola.com,1201804,DIRECT,c228e6794e811952
taboola.com,1201800,DIRECT,c228e6794e811952
taboola.com,1201822,DIRECT,c228e6794e811952
taboola.com,1201759,DIRECT,c228e6794e811952
taboola.com,1201806,DIRECT,c228e6794e811952
taboola.com,1201801,DIRECT,c228e6794e811952
taboola.com,1201798,DIRECT,c228e6794e811952
taboola.com,1201766,DIRECT,c228e6794e811952
taboola.com,1201819,DIRECT,c228e6794e811952
taboola.com,1201762,DIRECT,c228e6794e811952
taboola.com,1201797,DIRECT,c228e6794e811952
taboola.com,1201799,DIRECT,c228e6794e811952
taboola.com,1201757,DIRECT,c228e6794e811952
taboola.com,1201823,DIRECT,c228e6794e811952
taboola.com,1201803,DIRECT,c228e6794e811952
spotx.tv,71451,RESELLER
spotxchange.com, 71451, RESELLER
advertising.com, 8603, RESELLER
pubmatic.com, 156307, RESELLER, 5d62403b186f2ace
appnexus.com, 3364, RESELLER
Indexexchange.com, 183756, RESELLER
contextweb.com, 560382, RESELLER
openx.com, 539154393, RESELLER
tremorhub.com, z87wm, RESELLER, 1a4e959a1b50034a
rubiconproject.com, 16698, RESELLER, 0bfd66d529a55807
freewheel.tv, 799841, RESELLER 
freewheel.tv, 799921, RESELLER 
aol.com, 53392, RESELLER 
rhythmone.com, 1166984029, RESELLER, a670c89d4a324e47
`)
