<?php

class Request
{
    static public function Call(string $url, array $data, array $header = []): array
    {
        $dataStr = json_encode($data);
        var_dump($dataStr);
        $ch = curl_init();
        curl_setopt($ch, CURLOPT_URL, $url);
        curl_setopt($ch, CURLOPT_RETURNTRANSFER, 1);
        $h = ['Content-Type: application/json; charset=utf-8'];
        if (count($header) > 0) {
            $h = array_merge($h, $header);
        }
        $h[] = 'Content-Length:' . strlen($dataStr);
        curl_setopt($ch, CURLOPT_HTTPHEADER, $h);
        // post数据
        curl_setopt($ch, CURLOPT_POST, 1);
        // post的变量
        curl_setopt($ch, CURLOPT_POSTFIELDS, $dataStr);
        curl_setopt($ch, CURLOPT_TIMEOUT, 10);
        curl_setopt($ch, CURLOPT_CONNECTTIMEOUT, 2);

        //https 请求
        if (strlen($url) > 5 && strtolower(substr($url, 0, 5)) == "https") {
            curl_setopt($ch, CURLOPT_SSL_VERIFYPEER, false);
            curl_setopt($ch, CURLOPT_SSL_VERIFYHOST, false);
        }
        $rst = curl_exec($ch);
        $errno = curl_errno($ch);
        $error_msg = curl_error($ch);
        curl_close($ch);
        return [$rst, $errno, $error_msg];
    }
}
