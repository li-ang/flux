// DO NOT EDIT: This file is autogenerated via the builtin command.

package table

import (
	ast "github.com/influxdata/flux/ast"
	runtime "github.com/influxdata/flux/runtime"
)

func init() {
	runtime.RegisterPackage(pkgAST)
}

var pkgAST = &ast.Package{
	BaseNode: ast.BaseNode{
		Comments: nil,
		Errors:   nil,
		Loc:      nil,
	},
	Files: []*ast.File{&ast.File{
		BaseNode: ast.BaseNode{
			Comments: nil,
			Errors:   nil,
			Loc: &ast.SourceLocation{
				End: ast.Position{
					Column: 13,
					Line:   7,
				},
				File:   "table.flux",
				Source: "package table\n\n\n// fill will ensure that all tables within this stream have at least\n// one row. If a table has no rows, one row will be created with null values\n// for every column not part of the group key.\nbuiltin fill",
				Start: ast.Position{
					Column: 1,
					Line:   1,
				},
			},
		},
		Body: []ast.Statement{&ast.BuiltinStatement{
			BaseNode: ast.BaseNode{
				Comments: []ast.Comment{ast.Comment{Text: "// fill will ensure that all tables within this stream have at least\n"}, ast.Comment{Text: "// one row. If a table has no rows, one row will be created with null values\n"}, ast.Comment{Text: "// for every column not part of the group key.\n"}},
				Errors:   nil,
				Loc: &ast.SourceLocation{
					End: ast.Position{
						Column: 13,
						Line:   7,
					},
					File:   "table.flux",
					Source: "builtin fill",
					Start: ast.Position{
						Column: 1,
						Line:   7,
					},
				},
			},
			Colon: nil,
			ID: &ast.Identifier{
				BaseNode: ast.BaseNode{
					Comments: nil,
					Errors:   nil,
					Loc: &ast.SourceLocation{
						End: ast.Position{
							Column: 13,
							Line:   7,
						},
						File:   "table.flux",
						Source: "fill",
						Start: ast.Position{
							Column: 9,
							Line:   7,
						},
					},
				},
				Name: "fill",
			},
			Ty: ast.TypeExpression{
				BaseNode: ast.BaseNode{
					Comments: nil,
					Errors:   nil,
					Loc: &ast.SourceLocation{
						End: ast.Position{
							Column: 54,
							Line:   7,
						},
						File:   "table.flux",
						Source: "(<-tables: [A]) => [A] where A: Record",
						Start: ast.Position{
							Column: 16,
							Line:   7,
						},
					},
				},
				Constraints: []*ast.TypeConstraint{&ast.TypeConstraint{
					BaseNode: ast.BaseNode{
						Comments: nil,
						Errors:   nil,
						Loc: &ast.SourceLocation{
							End: ast.Position{
								Column: 54,
								Line:   7,
							},
							File:   "table.flux",
							Source: "A: Record",
							Start: ast.Position{
								Column: 45,
								Line:   7,
							},
						},
					},
					Kinds: []*ast.Identifier{&ast.Identifier{
						BaseNode: ast.BaseNode{
							Comments: nil,
							Errors:   nil,
							Loc: &ast.SourceLocation{
								End: ast.Position{
									Column: 54,
									Line:   7,
								},
								File:   "table.flux",
								Source: "Record",
								Start: ast.Position{
									Column: 48,
									Line:   7,
								},
							},
						},
						Name: "Record",
					}},
					Tvar: &ast.Identifier{
						BaseNode: ast.BaseNode{
							Comments: nil,
							Errors:   nil,
							Loc: &ast.SourceLocation{
								End: ast.Position{
									Column: 46,
									Line:   7,
								},
								File:   "table.flux",
								Source: "A",
								Start: ast.Position{
									Column: 45,
									Line:   7,
								},
							},
						},
						Name: "A",
					},
				}},
				Ty: &ast.FunctionType{
					BaseNode: ast.BaseNode{
						Comments: nil,
						Errors:   nil,
						Loc: &ast.SourceLocation{
							End: ast.Position{
								Column: 38,
								Line:   7,
							},
							File:   "table.flux",
							Source: "(<-tables: [A]) => [A]",
							Start: ast.Position{
								Column: 16,
								Line:   7,
							},
						},
					},
					Parameters: []*ast.ParameterType{&ast.ParameterType{
						BaseNode: ast.BaseNode{
							Comments: nil,
							Errors:   nil,
							Loc: &ast.SourceLocation{
								End: ast.Position{
									Column: 30,
									Line:   7,
								},
								File:   "table.flux",
								Source: "<-tables: [A]",
								Start: ast.Position{
									Column: 17,
									Line:   7,
								},
							},
						},
						Kind: "Pipe",
						Name: &ast.Identifier{
							BaseNode: ast.BaseNode{
								Comments: nil,
								Errors:   nil,
								Loc: &ast.SourceLocation{
									End: ast.Position{
										Column: 25,
										Line:   7,
									},
									File:   "table.flux",
									Source: "tables",
									Start: ast.Position{
										Column: 19,
										Line:   7,
									},
								},
							},
							Name: "tables",
						},
						Ty: &ast.ArrayType{
							BaseNode: ast.BaseNode{
								Comments: nil,
								Errors:   nil,
								Loc: &ast.SourceLocation{
									End: ast.Position{
										Column: 30,
										Line:   7,
									},
									File:   "table.flux",
									Source: "[A]",
									Start: ast.Position{
										Column: 27,
										Line:   7,
									},
								},
							},
							ElementType: &ast.TvarType{
								BaseNode: ast.BaseNode{
									Comments: nil,
									Errors:   nil,
									Loc: &ast.SourceLocation{
										End: ast.Position{
											Column: 29,
											Line:   7,
										},
										File:   "table.flux",
										Source: "A",
										Start: ast.Position{
											Column: 28,
											Line:   7,
										},
									},
								},
								ID: &ast.Identifier{
									BaseNode: ast.BaseNode{
										Comments: nil,
										Errors:   nil,
										Loc: &ast.SourceLocation{
											End: ast.Position{
												Column: 29,
												Line:   7,
											},
											File:   "table.flux",
											Source: "A",
											Start: ast.Position{
												Column: 28,
												Line:   7,
											},
										},
									},
									Name: "A",
								},
							},
						},
					}},
					Return: &ast.ArrayType{
						BaseNode: ast.BaseNode{
							Comments: nil,
							Errors:   nil,
							Loc: &ast.SourceLocation{
								End: ast.Position{
									Column: 38,
									Line:   7,
								},
								File:   "table.flux",
								Source: "[A]",
								Start: ast.Position{
									Column: 35,
									Line:   7,
								},
							},
						},
						ElementType: &ast.TvarType{
							BaseNode: ast.BaseNode{
								Comments: nil,
								Errors:   nil,
								Loc: &ast.SourceLocation{
									End: ast.Position{
										Column: 37,
										Line:   7,
									},
									File:   "table.flux",
									Source: "A",
									Start: ast.Position{
										Column: 36,
										Line:   7,
									},
								},
							},
							ID: &ast.Identifier{
								BaseNode: ast.BaseNode{
									Comments: nil,
									Errors:   nil,
									Loc: &ast.SourceLocation{
										End: ast.Position{
											Column: 37,
											Line:   7,
										},
										File:   "table.flux",
										Source: "A",
										Start: ast.Position{
											Column: 36,
											Line:   7,
										},
									},
								},
								Name: "A",
							},
						},
					},
				},
			},
		}},
		Eof:      nil,
		Imports:  nil,
		Metadata: "parser-type=rust",
		Name:     "table.flux",
		Package: &ast.PackageClause{
			BaseNode: ast.BaseNode{
				Comments: nil,
				Errors:   nil,
				Loc: &ast.SourceLocation{
					End: ast.Position{
						Column: 14,
						Line:   1,
					},
					File:   "table.flux",
					Source: "package table",
					Start: ast.Position{
						Column: 1,
						Line:   1,
					},
				},
			},
			Name: &ast.Identifier{
				BaseNode: ast.BaseNode{
					Comments: nil,
					Errors:   nil,
					Loc: &ast.SourceLocation{
						End: ast.Position{
							Column: 14,
							Line:   1,
						},
						File:   "table.flux",
						Source: "table",
						Start: ast.Position{
							Column: 9,
							Line:   1,
						},
					},
				},
				Name: "table",
			},
		},
	}},
	Package: "table",
	Path:    "experimental/table",
}
