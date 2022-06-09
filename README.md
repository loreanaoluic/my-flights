# My flights
<!-- TABLE OF CONTENTS -->
<details open="open">
  <summary>Sadržaj</summary>
  <ol>
    <li>
      <a href="#opis-aplikacije">Opis aplikacije</a>
    </li>
        <li>
      <a href="#funkcionalnosti-korisnika-aplikacije">Funkcionalnosti korisnika aplikacije</a>
    </li>
    <li>
      <a href="#arhitektura-sistema">Arhitektura sistema</a>
    </li>
    <li>
      <a href="#diplomski-rad">Diplomski rad</a>
    </li>
  </ol>
</details>


<!-- OPIS APLIKACIJE -->
## Opis aplikacije
My flights je aplikacija pomoću koje se vrši pretraživanje i rezervacija avionskih karata.

<!-- KORISNICI APLIKACIJE -->
## Funkcionalnosti korisnika aplikacije
* <b> Neautentifikovan korisnik: </b> <br>
  * PREGLED SVIH LETOVA - Omogućen mu je prikaz svih letova i njihova pretraga po kriterijumima kao što su relacija, datum, cena, avio kompanija, status (aktivan, nema dovoljno mesta, otkazan). Prikaz leta podrazumeva nazive mesta odlaska i dolaska, datum i vreme odlaska i dolaska, cenu, naziv avio kompanije, status.
  * PRIJAVA
  * REGISTRACIJA

<br>

* <b> Autentifikovan korisnik: </b> <br>
  * REZERVACIJA KARATA - Kada odabere let iz liste ponuđenih, korisniku je omogućena rezervacija istog (samo ukoliko je status leta aktivan) i tada bira broj putnika. Nakon uspešne rezervacije, korisnik ostvaruje određen broj poena koje će kasnije moći da iskoristi kao popust na narednu rezervaciju. Broj poena koje će ostvariti rezervacijom će biti prikazan u detaljnom prikazu leta pre same potvrde rezervacije, a skala na osnovu koje će se računati popust biće jasno definisana i prikazana. Takođe, dobija potvrdu rezervacije na email.
  * PREGLED SVIH AVIO-KOMPANIJA - Omogućen mu je prikaz svih avio-kompanija, ostavljenih komentara i ocena o njima. Nakon obavljenog leta, korisnik može ostaviti recenziju (komentar) o avio-kompaniji, kao i da je oceni. Omogućena mu je i prijava nedoličnog ponašanja (neprikladni komentari).
  * UVID U PROFIL I IZMENA ISTOG - Korisnik može izmeniti svoje lične informacije - ime, prezime, email adresu. Pored toga dat mu je uvid i u broj osvojenih poena. Ukoliko želi, korisnik može deaktivirati svoj profil. 

<br>

* <b> Admin: </b> <br>
  * CRUD NAD LETOVIMA - Pregled svih letova, dodavanje novih, kao i otkazivanje postojećih.
  * CRUD NAD AVIO-KOMPANIJAMA - Pregled svih avio-kompanija, dodavanje novih, kao i brisanje postojećih.
  * CRUD NAD KORISNICIMA - Pregled svih korisnika i njihovih informacija (ličnih, rezervisanih karata i broj ostvarenih poena). Može sortirati korisnike po tome koliko su sumnjivi (koliko su puta korisniku prijavljeni komentari), pri čemu će biti omogućeno banovanje korisnika na određeni vremenski period zbog nedoličnih komentara.

<!-- ARHITEKTURA SISTEMA -->
## Arhitektura sistema
* API Gateway: Go
* Mikroservis za korisnike (CRUD sa korisnicima) - Go
* Mikroservis za email (slanje potvrde o rezervaciji leta na email) - Go
* Mikroservis za letove (CRUD sa letovima) - Go
* Mikroservis za avio-kompanije (CRUD sa avio-kompanijama) - Go
* Mikroservis za rezervacije letova (servis za rezervisanje letova, poeni i popusti nakon rezervacije) - Go
* Mikroservis za komentare i ocene (servis za rad sa komentarima i ocenama) - Rust
* Klijentska veb aplikacija - Angular
* Baza podataka: PostgreSQL

<!-- DIPLOMSKI RAD -->
## Diplomski rad
Implementiranje naprednog algoritma za pretragu (Depth First Search) koji će korisniku omogućiti da prilikom pretrage leta prikaže sve moguće putanje između dva uneta aerodroma. Rezultati će biti rangirani po određenim kriterijumima (najjeftiniji, najmanje presedanja, najkraće vreme trajanja letova, itd.)

SW-60/2018 Loreana Oluić
