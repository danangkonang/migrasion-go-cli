package seed

import (
	"fmt"
	"log"

	"github.com/danangkonang/migrasion-go-cli/config"
)

func Kecamatan19to21() {
	db := config.Connect()
	_, err := db.Exec(`INSERT INTO kecamatan (kecamatan,kabupaten_id,provinsi_id)VALUES
		('DUSUN HILIR', 274, 19),
		('DUSUN SELATAN', 274, 19),
		('DUSUN UTARA', 274, 19),
		('GUNUNG BINTANG AWAI', 274, 19),
		('JENAMAS', 274, 19),
		('KARAU KUALA', 274, 19),
		('AWANG', 275, 19),
		('BENUA LIMA', 275, 19),
		('DUSUN TENGAH', 275, 19),
		('DUSUN TIMUR', 275, 19),
		('KARUSEN JANANG', 275, 19),
		('PAJU EPAT', 275, 19),
		('PAKU', 275, 19),
		('PATANGKEP TUTUI', 275, 19),
		('PEMATANG KARAU', 275, 19),
		('RAREN BATUAH', 275, 19),
		('GUNUNG PUREI', 276, 19),
		('GUNUNG TIMANG', 276, 19),
		('LAHEI', 276, 19),
		('LAHEI BARAT', 276, 19),
		('MONTALLAT (MONTALAT)', 276, 19),
		('TEWEH BARU', 276, 19),
		('TEWEH SELATAN', 276, 19),
		('TEWEH TENGAH', 276, 19),
		('TEWEH TIMUR', 276, 19),
		('DAMANG BATU', 277, 19),
		('KAHAYAN HULU UTARA', 277, 19),
		('KURUN', 277, 19),
		('MANUHING', 277, 19),
		('MANUHING RAYA', 277, 19),
		('MIHING RAYA', 277, 19),
		('MIRI MANASA', 277, 19),
		('RUNGAN', 277, 19),
		('RUNGAN BARAT', 277, 19),
		('RUNGAN HULU', 277, 19),
		('SEPANG (SEPANG SIMIN)', 277, 19),
		('TEWAH', 277, 19),
		('BASARANG', 278, 19),
		('BATAGUH', 278, 19),
		('DADAHUP', 278, 19),
		('KAPUAS BARAT', 278, 19),
		('KAPUAS HILIR', 278, 19),
		('KAPUAS HULU', 278, 19),
		('KAPUAS KUALA', 278, 19),
		('KAPUAS MURUNG', 278, 19),
		('KAPUAS TENGAH', 278, 19),
		('KAPUAS TIMUR', 278, 19),
		('MANDAU TALAWANG', 278, 19),
		('MANTANGAI', 278, 19),
		('PASAK TALAWANG', 278, 19),
		('PULAU PETAK', 278, 19),
		('TAMBAN CATUR', 278, 19),
		('TIMPAH', 278, 19),
		('BUKIT RAYA', 279, 19),
		('KAMIPANG', 279, 19),
		('KATINGAN HILIR', 279, 19),
		('KATINGAN HULU', 279, 19),
		('KATINGAN KUALA', 279, 19),
		('KATINGAN TENGAH', 279, 19),
		('MARIKIT', 279, 19),
		('MENDAWAI', 279, 19),
		('PETAK MALAI', 279, 19),
		('PULAU MALAN', 279, 19),
		('SANAMAN MANTIKEI (SENAMANG MANTIKEI)', 279, 19),
		('TASIK PAYAWAN', 279, 19),
		('TEWANG SANGGALANG GARING (SANGALANG)', 279, 19),
		('ARUT SELATAN', 280, 19),
		('ARUT UTARA', 280, 19),
		('KOTAWARINGIN LAMA', 280, 19),
		('KUMAI', 280, 19),
		('PANGKALAN BANTENG', 280, 19),
		('PANGKALAN LADA', 280, 19),
		('ANTANG KALANG', 281, 19),
		('BAAMANG', 281, 19),
		('BUKIT SANTUEI', 281, 19),
		('CEMPAGA', 281, 19),
		('CEMPAGA HULU', 281, 19),
		('KOTA BESI', 281, 19),
		('MENTAWA BARU (KETAPANG)', 281, 19),
		('MENTAYA HILIR SELATAN', 281, 19),
		('MENTAYA HILIR UTARA', 281, 19),
		('MENTAYA HULU', 281, 19),
		('PARENGGEAN', 281, 19),
		('PULAU HANAUT', 281, 19),
		('SERANAU', 281, 19),
		('TELAGA ANTANG', 281, 19),
		('TELAWANG', 281, 19),
		('TELUK SAMPIT', 281, 19),
		('TUALAN HULU', 281, 19),
		('BATANGKAWA', 282, 19),
		('BELANTIKAN RAYA', 282, 19),
		('BULIK', 282, 19),
		('BULIK TIMUR', 282, 19),
		('DELANG', 282, 19),
		('LAMANDAU', 282, 19),
		('MENTHOBI RAYA', 282, 19),
		('SEMATU JAYA', 282, 19),
		('BARITO TUHUP RAYA', 283, 19),
		('LAUNG TUHUP', 283, 19),
		('MURUNG', 283, 19),
		('PERMATA INTAN', 283, 19),
		('SERIBU RIAM', 283, 19),
		('SUMBER BARITO', 283, 19),
		('SUNGAI BABUAT', 283, 19),
		('TANAH SIANG', 283, 19),
		('TANAH SIANG SELATAN', 283, 19),
		('UUT MURUNG', 283, 19),
		('BUKIT BATU', 284, 19),
		('JEKAN RAYA', 284, 19),
		('PAHANDUT', 284, 19),
		('RAKUMPIT', 284, 19),
		('SEBANGAU', 284, 19),
		('BANAMA TINGANG', 285, 19),
		('JABIREN RAYA', 285, 19),
		('KAHAYAN HILIR', 285, 19),
		('KAHAYAN KUALA', 285, 19),
		('KAHAYAN TENGAH', 285, 19),
		('MALIKU', 285, 19),
		('PANDIH BATU', 285, 19),
		('SEBANGAU KUALA', 285, 19),
		('DANAU SELULUK', 286, 19),
		('DANAU SEMBULUH', 286, 19),
		('HANAU', 286, 19),
		('SERUYAN HILIR', 286, 19),
		('SERUYAN HILIR TIMUR', 286, 19),
		('SERUYAN HULU', 286, 19),
		('SERUYAN RAYA', 286, 19),
		('SERUYAN TENGAH', 286, 19),
		('SULING TAMBUN', 286, 19),
		('BALAI RIAM', 287, 19),
		('JELAI', 287, 19),
		('PANTAI LUNCI', 287, 19),
		('PERMATA KECUBUNG', 287, 19),
		('SUKAMARA', 287, 19),
		('AWAYAN', 288, 20),
		('BATU MANDI', 288, 20),
		('HALONG', 288, 20),
		('JUAI', 288, 20),
		('LAMPIHONG', 288, 20),
		('PARINGIN', 288, 20),
		('PARINGIN SELATAN', 288, 20),
		('TEBING TINGGI', 288, 20),
		('BANJAR BARU SELATAN', 290, 20),
		('BANJAR BARU UTARA', 290, 20),
		('CEMPAKA', 290, 20),
		('LANDASAN ULIN', 290, 20),
		('LIANG ANGGANG', 290, 20),
		('BANJARMASIN BARAT', 291, 20),
		('BANJARMASIN SELATAN', 291, 20),
		('BANJARMASIN TENGAH', 291, 20),
		('BANJARMASIN TIMUR', 291, 20),
		('BANJARMASIN UTARA', 291, 20),
		('ALALAK', 292, 20),
		('ANJIR MUARA', 292, 20),
		('ANJIR PASAR', 292, 20),
		('BAKUMPAI', 292, 20),
		('BARAMBAI', 292, 20),
		('BELAWANG', 292, 20),
		('CERBON', 292, 20),
		('JEJANGKIT', 292, 20),
		('MANDASTANA', 292, 20),
		('MARABAHAN', 292, 20),
		('MEKAR SARI', 292, 20),
		('RANTAU BADAUH', 292, 20),
		('TABUKAN', 292, 20),
		('TABUNGANEN', 292, 20),
		('TAMBAN', 292, 20),
		('WANARAYA', 292, 20),
		('ANGKINANG', 293, 20),
		('DAHA BARAT', 293, 20),
		('DAHA SELATAN', 293, 20),
		('DAHA UTARA', 293, 20),
		('KALUMPANG (KELUMPANG)', 293, 20),
		('KANDANGAN', 293, 20),
		('LOKSADO', 293, 20),
		('PADANG BATUNG', 293, 20),
		('SIMPUR', 293, 20),
		('TELAGA LANGSAT', 293, 20),
		('BARABAI', 294, 20),
		('BATANG ALAI SELATAN', 294, 20),
		('BATANG ALAI TIMUR', 294, 20),
		('BATANG ALAI UTARA', 294, 20),
		('BATU BENAWA', 294, 20),
		('HANTAKAN', 294, 20),
		('HARUYAN', 294, 20),
		('LABUAN AMAS SELATAN', 294, 20),
		('LABUAN AMAS UTARA', 294, 20),
		('LIMPASU', 294, 20),
		('PANDAWAN', 294, 20),
		('AMUNTAI SELATAN', 295, 20),
		('AMUNTAI TENGAH', 295, 20),
		('AMUNTAI UTARA', 295, 20),
		('BABIRIK', 295, 20),
		('BANJANG', 295, 20),
		('DANAU PANGGANG', 295, 20),
		('HAUR GADING', 295, 20),
		('PAMINGGIR', 295, 20),
		('SUNGAI PANDAN', 295, 20),
		('SUNGAI TABUKAN', 295, 20),
		('HAMPANG', 296, 20),
		('KELUMPANG BARAT', 296, 20),
		('KELUMPANG HILIR', 296, 20),
		('KELUMPANG HULU', 296, 20),
		('KELUMPANG SELATAN', 296, 20),
		('KELUMPANG TENGAH', 296, 20),
		('KELUMPANG UTARA', 296, 20),
		('PAMUKAN BARAT', 296, 20),
		('PAMUKAN SELATAN', 296, 20),
		('PAMUKAN UTARA', 296, 20),
		('PULAU LAUT BARAT', 296, 20),
		('PULAU LAUT KEPULAUAN', 296, 20),
		('PULAU LAUT SELATAN', 296, 20),
		('PULAU LAUT TANJUNG SELAYAR', 296, 20),
		('PULAU LAUT TENGAH', 296, 20),
		('PULAU LAUT TIMUR', 296, 20),
		('PULAU LAUT UTARA', 296, 20),
		('PULAU SEBUKU', 296, 20),
		('SAMPANAHAN', 296, 20),
		('SUNGAI DURIAN', 296, 20),
		('BANUA LAWAS', 297, 20),
		('BINTANG ARA', 297, 20),
		('HARUAI', 297, 20),
		('JARO', 297, 20),
		('KELUA (KLUA)', 297, 20),
		('MUARA HARUS', 297, 20),
		('MUARA UYA', 297, 20),
		('MURUNG PUDAK', 297, 20),
		('PUGAAN', 297, 20),
		('TANTA', 297, 20),
		('UPAU', 297, 20),
		('ANGSANA', 298, 20),
		('BATULICIN', 298, 20),
		('KARANG BINTANG', 298, 20),
		('KURANJI', 298, 20),
		('KUSAN HILIR', 298, 20),
		('KUSAN HULU', 298, 20),
		('MANTEWE', 298, 20),
		('SATUI', 298, 20),
		('SIMPANG EMPAT', 298, 20),
		('SUNGAI LOBAN', 298, 20),
		('BAJUIN', 299, 20),
		('BATI-BATI', 299, 20),
		('BUMI MAKMUR', 299, 20),
		('JORONG', 299, 20),
		('KINTAP', 299, 20),
		('KURAU', 299, 20),
		('PANYIPATAN', 299, 20),
		('PELAIHARI', 299, 20),
		('TAKISUNG', 299, 20),
		('TAMBANG ULANG', 299, 20),
		('BAKARANGAN', 300, 20),
		('BUNGUR', 300, 20),
		('CANDI LARAS SELATAN', 300, 20),
		('CANDI LARAS UTARA', 300, 20),
		('HATUNGUN', 300, 20),
		('LOKPAIKAT', 300, 20),
		('PIANI', 300, 20),
		('SALAM BABARIS', 300, 20),
		('TAPIN SELATAN', 300, 20),
		('TAPIN TENGAH', 300, 20),
		('TAPIN UTARA', 300, 20),
		('BALIKPAPAN BARAT', 301, 21),
		('BALIKPAPAN KOTA', 301, 21),
		('BALIKPAPAN SELATAN', 301, 21),
		('BALIKPAPAN TENGAH', 301, 21),
		('BALIKPAPAN TIMUR', 301, 21),
		('BALIKPAPAN UTARA', 301, 21),
		('BIATAN', 302, 21),
		('BIDUK-BIDUK', 302, 21),
		('DERAWAN (PULAU DERAWAN)', 302, 21),
		('GUNUNG TABUR', 302, 21),
		('KELAY', 302, 21),
		('MARATUA', 302, 21),
		('SAMBALIUNG', 302, 21),
		('SEGAH', 302, 21),
		('TABALAR', 302, 21),
		('TALISAYAN', 302, 21),
		('TANJUNG REDEB', 302, 21),
		('TELUK BAYUR', 302, 21),
		('BONTANG BARAT', 303, 21),
		('BONTANG SELATAN', 303, 21),
		('BONTANG UTARA', 303, 21),
		('BARONG TONGKOK', 304, 21),
		('BENTIAN BESAR', 304, 21),
		('BONGAN', 304, 21),
		('DAMAI', 304, 21),
		('JEMPANG', 304, 21),
		('LAHAM', 304, 21),
		('LINGGANG BIGUNG', 304, 21),
		('LONG APARI', 304, 21),
		('LONG BAGUN', 304, 21),
		('LONG HUBUNG', 304, 21),
		('LONG IRAM', 304, 21),
		('LONG PAHANGAI', 304, 21),
		('MANOR BULATIN (MOOK MANAAR BULATN)', 304, 21),
		('MELAK', 304, 21),
		('MUARA LAWA', 304, 21),
		('MUARA PAHU', 304, 21),
		('NYUATAN', 304, 21),
		('PENYINGGAHAN', 304, 21),
		('SEKOLAQ DARAT', 304, 21),
		('SILUQ NGURAI', 304, 21),
		('TERING', 304, 21),
		('ANGGANA', 305, 21),
		('KEMBANG JANGGUT', 305, 21),
		('KENOHAN', 305, 21),
		('KOTA BANGUN', 305, 21),
		('LOA JANAN', 305, 21),
		('LOA KULU', 305, 21),
		('MARANG KAYU', 305, 21),
		('MUARA BADAK', 305, 21),
		('MUARA JAWA', 305, 21),
		('MUARA KAMAN', 305, 21),
		('MUARA MUNTAI', 305, 21),
		('MUARA WIS', 305, 21),
		('SAMBOJA (SEMBOJA)', 305, 21),
		('SANGA-SANGA', 305, 21),
		('SEBULU', 305, 21),
		('TENGGARONG', 305, 21),
		('TENGGARONG SEBERANG', 305, 21),
		('BENGALON', 306, 21),
		('BUSANG', 306, 21),
		('KALIORANG', 306, 21),
		('KARANGAN', 306, 21),
		('KAUBUN', 306, 21),
		('KONGBENG', 306, 21),
		('LONG MESANGAT (MESENGAT)', 306, 21),
		('MUARA ANCALONG', 306, 21),
		('MUARA BENGKAL', 306, 21),
		('MUARA WAHAU', 306, 21),
		('RANTAU PULUNG', 306, 21),
		('SANDARAN', 306, 21),
		('SANGATTA SELATAN', 306, 21),
		('SANGATTA UTARA', 306, 21),
		('SANGKULIRANG', 306, 21),
		('TELEN', 306, 21),
		('TELUK PANDAN', 306, 21),
		('BATU ENGAU', 307, 21),
		('BATU SOPANG', 307, 21),
		('KUARO', 307, 21),
		('LONG IKIS', 307, 21),
		('LONG KALI', 307, 21),
		('MUARA KOMAM', 307, 21),
		('MUARA SAMU', 307, 21),
		('PASIR BELENGKONG', 307, 21),
		('TANAH GROGOT', 307, 21),
		('TANJUNG HARAPAN', 307, 21),
		('BABULU', 308, 21),
		('PENAJAM', 308, 21),
		('SEPAKU', 308, 21),
		('WARU', 308, 21),
		('LOA JANAN ILIR', 309, 21),
		('PALARAN', 309, 21),
		('SAMARINDA ILIR', 309, 21),
		('SAMARINDA KOTA', 309, 21),
		('SAMARINDA SEBERANG', 309, 21),
		('SAMARINDA ULU', 309, 21),
		('SAMARINDA UTARA', 309, 21),
		('SAMBUTAN', 309, 21),
		('SUNGAI KUNJANG', 309, 21),
		('SUNGAI PINANG', 309, 21)
	`)
	if err != nil {
		log.Fatal(err)
		fmt.Print("kecamatan not 19-21 insert")
	}
}