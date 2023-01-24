// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// DeviceInfoColumns holds the columns for the "device_info" table.
	DeviceInfoColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "type", Type: field.TypeString, Size: 2147483647},
		{Name: "os", Type: field.TypeString, Size: 2147483647},
		{Name: "browser", Type: field.TypeString, Size: 2147483647},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// DeviceInfoTable holds the schema information for the "device_info" table.
	DeviceInfoTable = &schema.Table{
		Name:       "device_info",
		Columns:    DeviceInfoColumns,
		PrimaryKey: []*schema.Column{DeviceInfoColumns[0]},
	}
	// GenerationsColumns holds the columns for the "generations" table.
	GenerationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "width", Type: field.TypeInt32},
		{Name: "height", Type: field.TypeInt32},
		{Name: "num_interference_steps", Type: field.TypeInt32},
		{Name: "guidance_scale", Type: field.TypeFloat32},
		{Name: "seed", Type: field.TypeInt},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"started", "succeeded", "failed", "rejected"}},
		{Name: "failure_reason", Type: field.TypeString, Size: 2147483647},
		{Name: "country_code", Type: field.TypeString, Size: 2147483647},
		{Name: "is_submitted_to_gallery", Type: field.TypeBool, Default: false},
		{Name: "is_public", Type: field.TypeBool, Default: false},
		{Name: "init_image_url", Type: field.TypeString, Size: 2147483647},
		{Name: "started_at", Type: field.TypeTime},
		{Name: "completed_at", Type: field.TypeTime},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "device_info_id", Type: field.TypeUUID},
		{Name: "model_id", Type: field.TypeUUID},
		{Name: "negative_prompt_id", Type: field.TypeUUID},
		{Name: "prompt_id", Type: field.TypeUUID},
		{Name: "scheduler_id", Type: field.TypeUUID},
		{Name: "user_id", Type: field.TypeUUID},
	}
	// GenerationsTable holds the schema information for the "generations" table.
	GenerationsTable = &schema.Table{
		Name:       "generations",
		Columns:    GenerationsColumns,
		PrimaryKey: []*schema.Column{GenerationsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "generations_device_info_generations",
				Columns:    []*schema.Column{GenerationsColumns[16]},
				RefColumns: []*schema.Column{DeviceInfoColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "generations_generation_models_generations",
				Columns:    []*schema.Column{GenerationsColumns[17]},
				RefColumns: []*schema.Column{GenerationModelsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "generations_negative_prompts_generations",
				Columns:    []*schema.Column{GenerationsColumns[18]},
				RefColumns: []*schema.Column{NegativePromptsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "generations_prompts_generations",
				Columns:    []*schema.Column{GenerationsColumns[19]},
				RefColumns: []*schema.Column{PromptsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "generations_schedulers_generations",
				Columns:    []*schema.Column{GenerationsColumns[20]},
				RefColumns: []*schema.Column{SchedulersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "generations_users_generations",
				Columns:    []*schema.Column{GenerationsColumns[21]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// GenerationModelsColumns holds the columns for the "generation_models" table.
	GenerationModelsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "name", Type: field.TypeString, Size: 2147483647},
		{Name: "is_free", Type: field.TypeBool, Default: false},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// GenerationModelsTable holds the schema information for the "generation_models" table.
	GenerationModelsTable = &schema.Table{
		Name:       "generation_models",
		Columns:    GenerationModelsColumns,
		PrimaryKey: []*schema.Column{GenerationModelsColumns[0]},
	}
	// GenerationOutputsColumns holds the columns for the "generation_outputs" table.
	GenerationOutputsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "image_url", Type: field.TypeString, Size: 2147483647},
		{Name: "upscaled_image_url", Type: field.TypeString, Size: 2147483647},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "generation_id", Type: field.TypeUUID},
	}
	// GenerationOutputsTable holds the schema information for the "generation_outputs" table.
	GenerationOutputsTable = &schema.Table{
		Name:       "generation_outputs",
		Columns:    GenerationOutputsColumns,
		PrimaryKey: []*schema.Column{GenerationOutputsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "generation_outputs_generations_generation_outputs",
				Columns:    []*schema.Column{GenerationOutputsColumns[5]},
				RefColumns: []*schema.Column{GenerationsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// NegativePromptsColumns holds the columns for the "negative_prompts" table.
	NegativePromptsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "text", Type: field.TypeString, Size: 2147483647},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// NegativePromptsTable holds the schema information for the "negative_prompts" table.
	NegativePromptsTable = &schema.Table{
		Name:       "negative_prompts",
		Columns:    NegativePromptsColumns,
		PrimaryKey: []*schema.Column{NegativePromptsColumns[0]},
	}
	// PromptsColumns holds the columns for the "prompts" table.
	PromptsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "text", Type: field.TypeString, Size: 2147483647},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// PromptsTable holds the schema information for the "prompts" table.
	PromptsTable = &schema.Table{
		Name:       "prompts",
		Columns:    PromptsColumns,
		PrimaryKey: []*schema.Column{PromptsColumns[0]},
	}
	// SchedulersColumns holds the columns for the "schedulers" table.
	SchedulersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "name", Type: field.TypeString, Size: 2147483647},
		{Name: "is_free", Type: field.TypeBool, Default: false},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// SchedulersTable holds the schema information for the "schedulers" table.
	SchedulersTable = &schema.Table{
		Name:       "schedulers",
		Columns:    SchedulersColumns,
		PrimaryKey: []*schema.Column{SchedulersColumns[0]},
	}
	// UpscalesColumns holds the columns for the "upscales" table.
	UpscalesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "width", Type: field.TypeInt32},
		{Name: "height", Type: field.TypeInt32},
		{Name: "scale", Type: field.TypeInt32},
		{Name: "duration_ms", Type: field.TypeInt32},
		{Name: "country_code", Type: field.TypeString, Size: 2147483647},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"started", "succeeded", "failed"}},
		{Name: "failure_reason", Type: field.TypeString, Size: 2147483647},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "device_info_id", Type: field.TypeUUID},
		{Name: "model_id", Type: field.TypeUUID},
		{Name: "user_id", Type: field.TypeUUID},
	}
	// UpscalesTable holds the schema information for the "upscales" table.
	UpscalesTable = &schema.Table{
		Name:       "upscales",
		Columns:    UpscalesColumns,
		PrimaryKey: []*schema.Column{UpscalesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "upscales_device_info_upscales",
				Columns:    []*schema.Column{UpscalesColumns[10]},
				RefColumns: []*schema.Column{DeviceInfoColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "upscales_upscale_models_upscales",
				Columns:    []*schema.Column{UpscalesColumns[11]},
				RefColumns: []*schema.Column{UpscaleModelsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "upscales_users_upscales",
				Columns:    []*schema.Column{UpscalesColumns[12]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// UpscaleModelsColumns holds the columns for the "upscale_models" table.
	UpscaleModelsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "name", Type: field.TypeString, Size: 2147483647},
		{Name: "is_free", Type: field.TypeBool, Default: false},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// UpscaleModelsTable holds the schema information for the "upscale_models" table.
	UpscaleModelsTable = &schema.Table{
		Name:       "upscale_models",
		Columns:    UpscaleModelsColumns,
		PrimaryKey: []*schema.Column{UpscaleModelsColumns[0]},
	}
	// UpscaleOutputsColumns holds the columns for the "upscale_outputs" table.
	UpscaleOutputsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "image_url", Type: field.TypeString, Size: 2147483647},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "upscale_id", Type: field.TypeUUID},
	}
	// UpscaleOutputsTable holds the schema information for the "upscale_outputs" table.
	UpscaleOutputsTable = &schema.Table{
		Name:       "upscale_outputs",
		Columns:    UpscaleOutputsColumns,
		PrimaryKey: []*schema.Column{UpscaleOutputsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "upscale_outputs_upscales_upscale_outputs",
				Columns:    []*schema.Column{UpscaleOutputsColumns[4]},
				RefColumns: []*schema.Column{UpscalesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "email", Type: field.TypeString, Size: 2147483647},
		{Name: "stripe_customer_id", Type: field.TypeString, Size: 2147483647},
		{Name: "subscription_category", Type: field.TypeEnum, Enums: []string{"GIFTED", "FRIEND_BOUGHT"}},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "confirmed_at", Type: field.TypeTime},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// UserRolesColumns holds the columns for the "user_roles" table.
	UserRolesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "role_name", Type: field.TypeEnum, Enums: []string{"ADMIN", "PRO"}},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "user_id", Type: field.TypeUUID},
	}
	// UserRolesTable holds the schema information for the "user_roles" table.
	UserRolesTable = &schema.Table{
		Name:       "user_roles",
		Columns:    UserRolesColumns,
		PrimaryKey: []*schema.Column{UserRolesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_roles_users_user_roles",
				Columns:    []*schema.Column{UserRolesColumns[4]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		DeviceInfoTable,
		GenerationsTable,
		GenerationModelsTable,
		GenerationOutputsTable,
		NegativePromptsTable,
		PromptsTable,
		SchedulersTable,
		UpscalesTable,
		UpscaleModelsTable,
		UpscaleOutputsTable,
		UsersTable,
		UserRolesTable,
	}
)

func init() {
	DeviceInfoTable.Annotation = &entsql.Annotation{
		Table: "device_info",
	}
	GenerationsTable.ForeignKeys[0].RefTable = DeviceInfoTable
	GenerationsTable.ForeignKeys[1].RefTable = GenerationModelsTable
	GenerationsTable.ForeignKeys[2].RefTable = NegativePromptsTable
	GenerationsTable.ForeignKeys[3].RefTable = PromptsTable
	GenerationsTable.ForeignKeys[4].RefTable = SchedulersTable
	GenerationsTable.ForeignKeys[5].RefTable = UsersTable
	GenerationsTable.Annotation = &entsql.Annotation{
		Table: "generations",
	}
	GenerationModelsTable.Annotation = &entsql.Annotation{
		Table: "generation_models",
	}
	GenerationOutputsTable.ForeignKeys[0].RefTable = GenerationsTable
	GenerationOutputsTable.Annotation = &entsql.Annotation{
		Table: "generation_outputs",
	}
	NegativePromptsTable.Annotation = &entsql.Annotation{
		Table: "negative_prompts",
	}
	PromptsTable.Annotation = &entsql.Annotation{
		Table: "prompts",
	}
	SchedulersTable.Annotation = &entsql.Annotation{
		Table: "schedulers",
	}
	UpscalesTable.ForeignKeys[0].RefTable = DeviceInfoTable
	UpscalesTable.ForeignKeys[1].RefTable = UpscaleModelsTable
	UpscalesTable.ForeignKeys[2].RefTable = UsersTable
	UpscalesTable.Annotation = &entsql.Annotation{
		Table: "upscales",
	}
	UpscaleModelsTable.Annotation = &entsql.Annotation{
		Table: "upscale_models",
	}
	UpscaleOutputsTable.ForeignKeys[0].RefTable = UpscalesTable
	UpscaleOutputsTable.Annotation = &entsql.Annotation{
		Table: "upscale_outputs",
	}
	UsersTable.Annotation = &entsql.Annotation{
		Table: "users",
	}
	UserRolesTable.ForeignKeys[0].RefTable = UsersTable
	UserRolesTable.Annotation = &entsql.Annotation{
		Table: "user_roles",
	}
}
