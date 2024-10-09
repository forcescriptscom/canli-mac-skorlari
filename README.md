# Canlı Maç Skorları API

**Canlı Maç Skorları API'si**, futbolseverler ve yazılım geliştiriciler için canlı futbol maç skorlarını, yaklaşan maçları ve popüler maç bilgilerini hızlı bir şekilde sunmaktadır. Futbol maçlarına dair detaylı bilgilere kolayca erişim sağlamak için ideal bir çözümdür. Bu API, futbol maçlarının anlık skorları, takım bilgileri, ülkeler, maç dakikası ve başlama saatleri gibi birçok veriyi sağlamaktadır.

## Özellikler

- **Canlı Maç Skorları**: Anlık güncellenen maç skorlarına erişim.
- **Takım ve Ülke Bilgileri**: Maç yapan takımların ve ülkelerin bilgilerini içerir.
- **Maç Dakikası ve Başlama Saati**: Maçların hangi dakikada olduğu ve başlama saatleri hakkında bilgi verir.
- **Yaklaşan Maçlar**: Gelecek maçlara dair detaylar.

## Kullanım Alanları

Bu API, canlı maç skorlarını görüntülemek isteyen web ve mobil uygulamalar için mükemmel bir çözümdür. Spor siteleri, mobil uygulamalar, futbol istatistik platformları gibi birçok farklı projede kullanılabilir.

## Satın Alma

Canlı maç skorları API'sini hızlı bir şekilde entegre etmek ve kullanmak için **[buraya tıklayarak](https://forcescripts.com/urun/futbol-maci-apisi-canli-maclar-yaklasan-maclar-ve-populer-maclar/)** satın alabilirsiniz. Bu bağlantıda, API'nin nasıl kullanılacağını gösteren örnek kodlar da mevcuttur.

## Fiyatlandırma

API, aylık ve 3 aylık abonelikler şeklinde sunulmaktadır. İhtiyacınıza uygun paketlerden birini seçerek kullanmaya başlayabilirsiniz.

## Örnek Kod

```php
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
```
