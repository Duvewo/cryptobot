package helpers

import (
	"bytes"
	"fmt"
	"github.com/Duvewo/cryptobot/internal/cryptocurrency"
	"github.com/gorilla/websocket"
	"log"
	"math/rand"
	"net/url"
	"strings"
)

//Generating Forex trading API URL to websocket and trying to dial
//Scheme: https://stream{server_id}.forexpros.com/echo/info
//WebSocket Scheme: wss://stream{server_id}.forexpros.com/echo/{rand 3 len int}/{random 8 len str}/websocket
func ForexDial() (*cryptocurrency.Client, error) {

	for i := 1; i <= 352; i++ {
		u, err := url.Parse(fmt.Sprintf("wss://stream%d.forexpros.com/echo/%d/%s/websocket", i, rand.Intn(999)+100, randomString(8)))

		if err != nil {
			return nil, err
		}

		cc, err := cryptocurrency.Dial(u)

		if err != nil {
			return nil, err
		}

		if cc.Ping() != nil {
			return cc, nil
		}

	}

	return nil, fmt.Errorf("crypto: failed to dial Forex")

}

//TODO: wrap errors
func ForexInit(client *cryptocurrency.Client) error {
	messageType, data, err := client.Read()

	if err != nil {
		return err
	}

	if messageType == websocket.TextMessage && bytes.Equal(data, []byte("o")) {
		log.Printf("Got success message from client!\n")
		err := client.WriteJSON(cryptocurrency.Request{Event: "bulk-subscribe", TzID: 18, Message: "pid-eu-1057391:%%pid-eu-1035793:%%pid-eu-1129220:%%pid-eu-945629:%%pid-eu-1062798:%%pid-eu-1057388:%%pid-eu-1062795:%%pid-eu-1064952:%%pid-eu-1059808:%%pid-eu-49798:%%pid-eu-1061443:%%pid-eu-1035794:%%pid-eu-1062799:%%pid-eu-1118047:%%pid-eu-1064958:%%pid-eu-1061972:%%pid-eu-1058142:%%pid-eu-997650:%%pid-eu-1170214:%%pid-eu-1068307:%%pid-eu-1061448:%%pid-eu-1054919:%%pid-eu-1068329:%%pid-eu-1036977:%%pid-eu-1072478:%%pid-eu-1036978:%%pid-eu-1170216:%%pid-eu-1136164:%%pid-eu-1136166:%%pid-eu-1161441:%%pid-eu-1062537:%%pid-eu-1073899:%%pid-eu-1064925:%%pid-eu-1064953:%%pid-eu-1141214:%%pid-eu-1068312:%%pid-eu-1055844:%%pid-eu-1141216:%%pid-eu-1064738:%%pid-eu-1072528:%%pid-eu-1061477:%%pid-eu-1158819:%%pid-eu-1128789:%%pid-eu-1068336:%%pid-eu-1167165:%%pid-eu-1128793:%%pid-eu-1173204:%%pid-eu-1169563:%%pid-eu-1173208:%%pid-eu-1128860:%%pid-eu-1061453:%%pid-eu-1129325:%%pid-eu-1031397:%%pid-eu-1128662:%%pid-eu-1128969:%%pid-eu-1161868:%%pid-eu-1097586:%%pid-eu-1056903:%%pid-eu-1065402:%%pid-eu-1056905:%%pid-eu-1057392:%%pid-eu-1062800:%%pid-eu-1064954:%%pid-eu-1075586:%%pid-eu-1058260:%%pid-eu-1068308:%%pid-eu-1054876:%%pid-eu-1118146:%%pid-eu-1062797:%%pid-eu-1031674:%%pid-eu-1165465:%%pid-eu-1169564:%%pid-eu-1169565:%%pid-eu-1166553:%%pid-eu-1168536:%%pid-eu-1166552:%%pid-eu-1168537:%%pid-eu-1173270:%%pid-eu-1173272:%%pid-eu-1173271:%%pid-eu-1061410:%%pid-eu-1114308:%%pid-eu-1099022:%%pid-eu-1058255:%%pid-eu-1114306:%%pid-eu-1064923:%%pid-eu-1064957:%%pid-eu-1068309:%%pid-eu-1099021:%%pid-eu-1114313:%%pid-eu-1167226:%%pid-eu-1167227:%%pid-eu-1171988:%%pid-eu-1171987:%%pid-eu-1170079:%%pid-eu-1061445:%%pid-eu-1073247:%%pid-eu-1056828:%%pid-eu-1117790:%%pid-eu-1058201:%%pid-eu-1057982:%%pid-eu-1064924:%%pid-eu-1064964:%%pid-eu-1068310:%%pid-eu-1061976:%%pid-eu-1061794:%%pid-eu-1137399:%%pid-eu-1137402:%%pid-eu-1068382:%%pid-eu-1056808:%%pid-eu-1057857:%%pid-eu-1072552:%%pid-eu-1056842:%%pid-eu-1137529:%%pid-eu-1137403:%%pid-eu-1061451:%%pid-eu-1122657:%%pid-eu-1064927:%%pid-eu-1064965:%%pid-eu-1057463:%%pid-eu-1129222:%%pid-eu-1122684:%%pid-eu-1122656:%%pid-eu-1068313:%%pid-eu-1057248:%%pid-eu-1114630:%%pid-eu-1142432:%%pid-eu-1142433:%%pid-eu-1099576:%%pid-eu-1142434:%%pid-eu-1099573:%%pid-eu-1114632:%%pid-eu-1099575:%%pid-eu-1114629:%%pid-eu-1099578:%%pid-eu-1161529:%%pid-eu-1161530:%%pid-eu-1161528:%%pid-eu-1161531:%%pid-eu-1161532:%%pid-eu-1057275:%%pid-eu-1129149:%%pid-eu-1064935:%%pid-eu-1064955:%%pid-eu-1129153:%%pid-eu-1122278:%%pid-eu-1057972:%%pid-eu-1031675:%%pid-eu-1068318:%%pid-eu-1058272:%%pid-eu-1061537:%%pid-eu-1089757:%%pid-eu-1068321:%%pid-eu-1089759:%%pid-eu-1056801:%%pid-eu-1072538:%%pid-eu-1056805:%%pid-eu-1056712:%%pid-eu-1056869:%%pid-eu-1089758:%%pid-eu-1138411:%%pid-eu-1131278:%%pid-eu-1131280:%%pid-eu-1138412:%%pid-eu-1131285:%%pid-eu-1173201:%%pid-eu-1173187:%%pid-eu-1131279:%%pid-eu-1131284:%%pid-eu-1066736:%%pid-eu-1158889:%%pid-eu-1062196:%%pid-eu-1115837:%%pid-eu-1062054:%%pid-eu-1068467:%%pid-eu-1097726:%%pid-eu-1062195:%%pid-eu-1072923:%%pid-eu-1115833:%%pid-eu-1061444:%%pid-eu-1129126:%%pid-eu-1058198:%%pid-eu-1064928:%%pid-eu-1119419:%%pid-eu-1119422:%%pid-eu-1119415:%%pid-eu-1119421:%%pid-eu-1064741:%%pid-eu-1156867:%%pid-eu-1061450:%%pid-eu-1116590:%%pid-eu-1116580:%%pid-eu-1116596:%%pid-eu-1064934:%%pid-eu-1123206:%%pid-eu-1057454:%%pid-eu-1064747:%%pid-eu-1068320:%%pid-eu-1054993:%%pid-eu-1172746:%%pid-eu-1167347:%%pid-eu-1167349:%%pid-eu-1167348:%%pid-eu-1061489:%%pid-eu-1172738:%%pid-eu-1057927:%%pid-eu-1172737:%%pid-eu-1172750:%%pid-eu-1068434:%%pid-eu-1056807:%%pid-eu-1057872:%%pid-eu-1061509:%%pid-eu-1173190:%%pid-eu-1072577:%%pid-eu-1056832:%%pid-eu-1173200:%%pid-eu-1058159:%%pid-eu-1170080:%%pid-eu-1162990:%%pid-eu-1162991:%%pid-eu-1173198:%%pid-eu-1173182:%%pid-eu-1173197:%%pid-eu-1173274:%%pid-eu-1061467:%%pid-eu-1057539:%%pid-eu-1068317:%%pid-eu-1031707:%%pid-eu-1054882:%%pid-eu-1057472:%%pid-eu-1061983:%%pid-eu-1024869:%%pid-eu-1072451:%%pid-eu-1031709:%%pid-eu-1061409:%%pid-eu-1121736:%%pid-eu-1064959:%%pid-eu-1064926:%%pid-eu-1055863:%%pid-eu-1057446:%%pid-eu-1068311:%%pid-eu-1064739:%%pid-eu-1031691:%%pid-eu-1162989:%%cmt-7-5-945629:%%domain-7:"})

		if err != nil {
			return err
		}

		err = client.WriteJSON(cryptocurrency.Request{Event: "UID", UID: 0})

		if err != nil {
			return err
		}

	}

	return nil

}

func randomString(n int) string {
	const s = "abcdefghijklmnopqrstuvwxyz0123456789_"

	var sb strings.Builder
	for i := 0; i < n; i++ {
		sb.WriteString(string(s[rand.Intn(len(s))]))
	}

	return sb.String()

}
