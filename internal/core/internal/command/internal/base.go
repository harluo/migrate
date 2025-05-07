package internal

import (
	"path/filepath"
	"reflect"

	"github.com/harluo/migrate/internal/kernel"
)

type Base struct{}

func (b *Base) Migrations(required kernel.Migration, optionals ...kernel.Migration) (migrations []map[string]any) {
	migrations = make([]map[string]any, 0, len(optionals)+1)
	for _, migration := range append([]kernel.Migration{required}, optionals...) {
		migrations = append(migrations, map[string]any{
			"filename":    b.sourceFilename(migration),
			"id":          migration.Id(),
			"description": migration.Description(),
		})
	}

	return
}

func (*Base) sourceFilename(target any) string {
	typeOf := reflect.TypeOf(target)
	if typeOf.Kind() == reflect.Ptr {
		typeOf = typeOf.Elem()
	}

	return filepath.Join(typeOf.PkgPath(), typeOf.Name()+".go")
}
