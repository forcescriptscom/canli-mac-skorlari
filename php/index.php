<?php
require 'vendor/autoload.php';

use GuzzleHttp\Client;

$client = new Client();
$api_url = 'https://apiv.forcescripts.com/DEMO';

try {
    // API'ye istek gönderme
    $response = $client->request('GET', $api_url);
    
    // Yanıtı JSON olarak çözme
    $data = json_decode($response->getBody(), true);

    // Canlı skorlar ve maç bilgilerini ekrana yazdırma
    foreach ($data['data'] as $match) {
        echo "Takımlar: " . $match['home_team'] . " vs " . $match['away_team'] . PHP_EOL;
        echo "Skor: " . $match['score']['homeTeam']['current'] . " - " . $match['score']['awayTeam']['current'] . PHP_EOL;
        echo "Dakika: " . $match['score']['minute'] . PHP_EOL;
        echo "Başlama Saati: " . date('Y-m-d H:i:s', $match['date']) . PHP_EOL;
        echo PHP_EOL; // Satır arası boşluk
    }
} catch (\Exception $e) {
    echo 'Hata: ' . $e->getMessage();
}
