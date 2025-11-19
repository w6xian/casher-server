<?php
 
class Rsa
{
    /**
     * 通过公钥加密，对应Go中的RsaEncrypt函数
     * @param string $plaintext 明文数据
     * @param string $publicKey 公钥字符串(PEM格式)
     * @return string|false 加密后base64编码的字符串，失败返回false
     */
    public static function RsaEncrypt(string $plaintext, string $publicKey): array
    {
        $encrypted = '';
// 使用公钥加密，采用PKCS#1 v1.5填充方式
        $result = openssl_public_encrypt($plaintext, $encrypted, $publicKey, OPENSSL_PKCS1_PADDING);
        if (!$result) {
            return ['', false];
        }
// 对加密结果进行base64编码
        return [base64_encode($encrypted), true];
    }

    /**
     * 通过私钥解密，对应Go中的RsaDecrypt函数
     * @param string $encrypted 加密后base64编码的字符串
     * @param string $privateKey 私钥字符串(PEM格式)
     * @return string|false 解密后的明文，失败返回false
     */
    public static function RsaDecrypt(string $encrypted, string $privateKey): array
    {
// 先进行base64解码
        $decoded = base64_decode($encrypted);
        if ($decoded === false) {
            return ['', false];
        }
        $decrypted = '';
// 使用私钥解密，采用PKCS#1 v1.5填充方式
        $result = openssl_private_decrypt($decoded, $decrypted, $privateKey, OPENSSL_PKCS1_PADDING);
        if (!$result) {
            return ['', false];
        }
        return [$decrypted, true];
    }
}