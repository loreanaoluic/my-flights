# My flights

<img alt="Pharo" src="https://img.shields.io/badge/pharo%20-%234f0599.svg?&style=for-the-badge"/> <img alt="Go" src="https://img.shields.io/badge/go-%2300ADD8.svg?&style=for-the-badge&logo=go&logoColor=white"/> <img alt="Pharo" src="https://img.shields.io/badge/postgres-%23316192.svg?style=for-the-badge&logo=postgresql&logoColor=white"/>

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
      <a href="#tehnologije">Tehnologije</a>
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
  * REZERVACIJA KARATA - Kada odabere let iz liste ponuđenih, korisniku je omogućena rezervacija istog (samo ukoliko je status leta aktivan) i tada bira broj putnika. Nakon uspešne rezervacije, korisnik ostvaruje određen broj poena koje će kasnije moći da iskoristi kao popust na narednu rezervaciju. Broj poena koje će ostvariti rezervacijom će biti prikazan u detaljnom prikazu leta pre same potvrde rezervacije, a skala na osnovu koje će se računati popust biće jasno definisana i prikazana.
  * UVID U PROFIL I IZMENA ISTOG - Korisnik može izmeniti svoje lične informacije - ime, prezime, email adresu. Pored toga dat mu je uvid i u broj osvojenih poena. Ukoliko želi, korisnik može deaktivirati svoj profil. 

<br>

* <b> Admin: </b> <br>
  * CRUD NAD LETOVIMA - Pregled svih letova, dodavanje novih, kao i otkazivanje postojećih.
  * CRUD NAD KORISNICIMA - Pregled svih korisnika i njihovih informacija (ličnih, rezervisanih karata i broj ostvarenih poena), banovanje korisnika, kao i ponovna aktivacija naloga nakon banovanja. Ukoliko admin odluči da banuje određenog korisnika, korisnik će o tome biti obavešten prilikom prvog narednog prijavljivanja na sistem koje će mu biti onemogućeno.

<!-- TEHNOLOGIJE -->
## Tehnologije
* Frontend: Pharo

* Backend: Golang

* Baza podataka: PostgreSQL


SW-60/2018 Loreana Oluić
