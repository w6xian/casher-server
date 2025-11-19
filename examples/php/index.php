<?php
$aesKey = <<<EOD
-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQC6B8yicpX5alZUQTuGimkRy2R3
rwwirRC3OJkL3Z+uzlBxiJm0EtTNd7QMk15xBKwfmDvvmeZ/vKf58v+6LJXR40W+
0/PoaW613XVeHGx8seq53QLi65OPkwfnlVTGK1mrjMf+GqMIjsNaMtWSP4nOtOkD
Q+VMScfbSQOt1tpFHwIDAQAB
-----END PUBLIC KEY-----
EOD;

$rd = $data['app_id'] . ':' . $data['ts'];
[$riskData, $b] = Rsa::RsaEncrypt($rd, $aesKey);
if (!$b) return;
$data['sign'] = $riskData;
var_dump(Request::Call("http://localhost:8966/system-info", $data));
exit;
