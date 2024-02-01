package integration

import (
	"bares_api/store"
	"fmt"
	"log"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestDatabaseIntegration(t *testing.T) {
  // Inicia o banco de dados
  dbStore, err := StartDatabase()
  if err != nil {
    t.Errorf("Falha ao criar banco de dados e tabelas: %s", err)
  }
  defer dbStore.DatabaseClose() // Garante que a conexão será fechada no final

  // Verificar a existência das tabelas
  tables := []string{
    store.TableUsuarios,
    store.TableItensMenu,
    store.TableItensPedido,
    store.TablePedidos,
  }
  for _, table := range tables {
    var tableName string
    query := fmt.Sprintf("SHOW TABLES LIKE '%s'", table)
    err := dbStore.DB.QueryRow(query).Scan(&tableName)
    if err != nil {
      t.Errorf("Tabela %s não foi criada corretamente", table)
    }
  }

  // Limpeza final
  if _, err := dbStore.DB.Exec(fmt.Sprintf("DROP DATABASE IF EXISTS %s", testDBName)); err != nil {
    log.Fatal("Falha ao limpar o banco de dados de teste após o teste:", err)
  }
}
