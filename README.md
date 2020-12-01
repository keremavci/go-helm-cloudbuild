Herkese merhaba, bu yazıda örnek bir uygulamayı Google CloudBuild kullanarak GKE üzerinde farklı ortamlara deploy etmekten bahsedeceğim.

Yazıda ele alacağımız örnek uygulama Golang ile yazılmış ve bir tane endpoint'i olan bir REST API. Örnek uygulamayı deploy edeceğimiz ortam ise Staging ve Production namespace'leri bulunan bir Kubernetes cluster'ı.

Uygulamamızın deployment senaryosu şu şekilde olacak;

1. Master branch'inde yapılan geliştirmeler release branch'ine merge edildiğinde, uygulama build edilip Staging namespace'ine deploy edilecek.
2. Her şey yolunda ise Production kelimesi ile başlayan, yeni oluşturulacak tag ile Staging namespace'ine deploy edilen versiyon Production namespace'ine de deploy edilecek.

Ben Kubernetes üzerine yapılacak deployment'larda Helm kullanmayı tercih ediyorum. Bu nedenle projede helm dizini altında çok basit bir chart oluşturdum. Burada  değişkenleri tanımladığım values.yml, values-staging.yml ve values.production.yml olarak 3 adet values dosyası mevcut. values.yml dosyasında uygulamamı deploy edeceğim tüm ortamlar için ortak olacak değerler varken values-staging.yml ve values-production.yml dosyamda ortamlara göre değişkenlik gösterecek değerler bulunuyor.

Projemizde Google CloudBuild Pipeline'ının adımlarını tanımladığımız cloudbuild.yaml ve cloudbuild-production.yaml adında iki tane dosyamız var. Bu iki dosyamıza karşılık Google Cloud Console'da bulunan CloudBuild Dashboard ekranında iki tane trigger oluşturacağız.

Bu işlem için önce Cloud Build ekranından sol menüden Triggers'a girdikten sonra Connect Repository butonuna tıklayarak GitHub Repolarımıza erişim izni vermemiz gerekiyor.

![Alt text](docs/1.jpg?raw=true "Connect Repository")

GitHub için gerekli izinleri verdikten sonra yine Triggers ekranında Create Trigger butonuna tıklayarak cloudbuild.yaml dosyamız için ilk triggerımızı oluşturacağız.

![Alt text](docs/2.jpg?raw=true "Create CloudBuild Staging Trigger ")

Burada şunları belirtiyoruz;
 - Trigger'ın adı. Bizim senaryomuzda build-and-deploy-to-staging olacak.
 - Event olarak daha önce bahsettiğim gibi release branch'ine yapılacak push ile pipeline tetiklenecek. Bunuda Event bölümünde "Push to a branch" seçeneği ve Source bölümünde Branch alanına release yazarak sağlıyoruz.
 - Son olarak Build configuration bölümünde Cloud Build configuration file (yaml or json) seçeneceğini seçerek triggerımızın kullanacağı cloudbuild.yaml dosyasını gösteriyoruz.

Triggers ekranında Create Trigger butonuna tıklayarak cloudbuild.yaml dosyamız için ilk triggerımızı oluşturacağız.


![Alt text](docs/3.jpg?raw=true "Create CloudBuild Production Trigger ")


Burada ise  şunları belirtiyoruz,
 - Trigger'ın adı. Bizim senaryomuzda bu kez deploy-to-production olacak.
 - Event olarak ise bu sefer Push new tag seçeneceğini seçiyoruz ve Source bölümündeki Tag kısmına ^production yazıyoruz.
 - Son olarak Build configuration bölümünde "Cloud Build configuration file (yaml or json)" seçeneceğini seçerek triggerımızın kullanacağı cloudbuild-production.yaml dosyasını gösteriyoruz.

Bu aşamalarla Google Cloud tarafındaki tanımlamalarımız bitti.

cloudbuild dosyalarımızı inceledeğimizde de temel olarak substitutions ve steps olarak iki bölümden oluşuyor. substitutions kısmında steps alanında çalışacak tasklarda kullanılacak değişkenleri tanımlıyoruz. steps alanında ise taskları tanımlıyoruz.

CloudBuild ile çalıştığımızda dikkat etmemiz gereken bir kaç nokta var. Bunlar şu şekilde,
 - İlk olarak steps alanında tanımlanan task'lar paralel şekilde çalışıyor. Bir task'ın bir diğer task'ı beklemesi ve sonrasında çalışması için waitFor ifadesinden yararlanıyoruz.
 - Task'ların içinde substitutions alanında tanımlanan değişkenleri kullanmak için $VAR_NAME şeklinde kullanılırken task içindeki ya da taskın çalışacağı container içindeki değişkenler $$VAR_NAME şeklinde kullanılıyor. 
 - CloudBuild defaultta çalışma dizini olarak /workspace'i kullanıyor.




Umarım faydalı olmuştur. Projeye https://github.com/keremavci/go-helm-cloudbuild adresinden erişebilirsiniz