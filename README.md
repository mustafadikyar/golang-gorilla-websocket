# golang-gorilla-websocket

Bu basit bir Go (Golang) WebSocket sunucusunu içeren bir projedir. 

Temel bir WebSocket sunucusu oluşturmayı ve WebSocket bağlantılarını nasıl yöneteceğinizi gösterir.

## Kullanım

1. Projeyi klonlayın:

```git
git clone https://github.com/mustafadikyar/golang-websocket.git
cd golang-websocket
```

2. Sunucuyu çalıştırın:

```shell
go run main.go
```

Sunucu `:8080` portunda dinlemeye başlayacaktır.

3. WebSocket sunucusuna bağlanın:
Bir WebSocket istemci veya tarayıcı uzantısı kullanarak `ws://localhost:8080/websocket` adresine bağlanabilirsiniz.

4. Mesaj gönderme ve alma:
Bağlandıktan sonra sunucuya mesajlar gönderebilirsiniz. Sunucu, aldığı mesajları size geri gönderecektir.

## WebSocket testi yapmak için Chrome'da "WebSocket Test Client" adlı bir uzantıyı kullanabilirsiniz. 

#### Uzantıyı Chrome'a Yükleyin:

- Chrome Web Mağazası'na gidin ve *"WebSocket Test Client"* uzantısını arayın.
Uzantıyı bulduktan sonra "Add to Chrome" (Chrome'a Ekle) düğmesine tıklayarak uzantıyı Chrome tarayıcınıza ekleyin.

#### Uzantıyı Açın:

- Uzantı eklendikten sonra, Chrome tarayıcınızda sağ üst köşede yeni bir simge göreceksiniz. Bu simgeye tıklayarak *"WebSocket Test Client"* uzantısını açın.

#### WebSocket Sunucusuna Bağlanın:

- "Server URL" (Sunucu URL'si) alanına test etmek istediğiniz WebSocket sunucusunun URL'sini girin. Bu proje için: **ws://localhost:8080/websocket**
Bağlantı yapmak için **"Connect"** (Bağlan) düğmesine tıklayın.

#### Mesaj Gönderme ve Alma:

- Bağlandıktan sonra, **"Message"** (Mesaj) alanına göndermek istediğiniz mesajı yazın.
Mesajınızı yazdıktan sonra **"Send"** (Gönder) düğmesine tıklayarak mesajı sunucuya gönderin.
Sunucudan gelen mesajlar **"Received Messages"** (Alınan Mesajlar) bölümünde görüntülenecektir.

#### Bağlantıyı Kapatma:

Test tamamlandığında veya bağlantıyı kapatmak istediğinizde **"Disconnect"** (Bağlantıyı Kapat) düğmesine tıklayarak WebSocket bağlantısını kapatabilirsiniz.
Bu adımları takip ederek, "WebSocket Test Client" uzantısı ile WebSocket sunucusuyla iletişim kurabilir ve mesajlar gönderebilirsiniz. Bu şekilde sunucu ve istemci arasındaki WebSocket iletişimini test edebilirsiniz.

## Bağımlılıklar

Bu proje, WebSocket bağlantılarını yönetmek için Gorilla WebSocket kütüphanesini kullanır. Gorilla WebSocket hakkında daha fazla bilgi için: [Gorilla WebSocket](https://github.com/gorilla/websocket).

