# twiscode

### 2. Palindrome

```go
package main

import (
	"fmt"
	"strings"
)

func palindrome(input string) bool {
	for i := 0; i < len(input)/2; i++ {
		if input[i] != input[len(input)-i-1] {
			return false
		}
	}
	return input
}

func main() {
    var n string
	fmt.Scan(&n)
	fmt.Println(palindrome(n))
}
```

## query test

### 1. 
```sql
CREATE TABLE TransactionData (
    ID int,
    TanggalPembayaran datetime,
    TanggalOrder datetime,
    IsLunas bool,
    PRIMARY KEY (ID),
);
```

```sql
INSERT INTO TransactionData(ID, TanggalPembayaran, TanggalOrder, IsLunas) values (1, 2021-03-09 21:26:26, 2021-03-09 21:26:26, true)
```

### 2.
```sql
CREATE TABLE DetilTransaksi (
    ID int,
    IDTransaksi int,
    Harga int,
    Jumlah int,
    Subtotal int
    PRIMARY KEY (ID),
    FOREIGN KEY (IDTransaksi) REFERENCES TransactionData(ID)
);
```

```sql
INSERT INTO DetilTransaksi(ID, IDTransaksi, Harga, Jumlah, Subtotal) values (1, 1, 20000, 3, 60000)
```

### 3

```sql
select dt.id, td.tanggal_order, td.status, td.tanggal_pembayaran, dt.total, dt.Jumlah as jumlah_barang 
from DetilTransaksi dt
join TransactionData td on td.id = dt.IDTransaksi 
```
