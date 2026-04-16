package antrian

// Entity untuk antrian pasien
// id: id antrian (auto increment)
// idPasien: id pasien yang dirujuk
// namaPasien: nama pasien
// idDokter: id dokter tujuan
// nomorAntrian: nomor urut antrian
// idPoli: id poli tujuan
// waktu: waktu masuk antrian

type Antrian struct {
	ID           int
	IDPasien     int
	IDDokter     int
	NomorAntrian int
	IDPoli       int
	Waktu        string // atau time.Time jika ingin parsing otomatis
}
