// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.793
package login

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"github.com/iota-uz/iota-sdk/components/base/alert"
	"github.com/iota-uz/iota-sdk/components/base/button"
	"github.com/iota-uz/iota-sdk/components/base/input"
	"github.com/iota-uz/iota-sdk/components/icons"
	"github.com/iota-uz/iota-sdk/internal/assets"
	"github.com/iota-uz/iota-sdk/modules/core/presentation/templates/layouts"
	"github.com/iota-uz/iota-sdk/pkg/types"
)

type LoginProps struct {
	*types.PageContext
	ErrorsMap          map[string]string
	ErrorMessage       string
	Email              string
	GoogleOAuthCodeURL string
}

var (
	navbarOnce  = templ.NewOnceHandle()
	companyLogo = "/assets/" + assets.HashFS.HashName("images/logo.webp")
)

func Header() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<section class=\"h-16 shadow-b-lg border-b w-full flex items-center justify-center px-8 bg-surface-300 border-b-primary\"><a href=\"/\"><img src=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(companyLogo)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `modules/core/presentation/templates/pages/login/index.templ`, Line: 29, Col: 25}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"h-10 w-auto\"></a><div class=\"ml-auto flex items-center gap-8\"><div class=\"flex items-center justify-center w-9 h-9\"><input class=\"peer/system appearance-none absolute\" type=\"radio\" name=\"theme\" value=\"system\" id=\"theme-system\" onchange=\"onThemeChange(this)\" checked> <label for=\"theme-light\" class=\"group/system absolute flex items-center justify-center w-9 h-9 rounded-full bg-gray-200 text-black invisible peer-checked/system:visible\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = icons.Desktop(icons.Props{Size: "20", Class: "scale-0 duration-200 peer-checked/system:group-[]/system:scale-100"}).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</label> <input class=\"peer/light appearance-none absolute\" type=\"radio\" name=\"theme\" value=\"light\" id=\"theme-light\" onchange=\"onThemeChange(this)\"> <label for=\"theme-dark\" class=\"group/light absolute flex items-center justify-center w-9 h-9 rounded-full bg-gray-200 text-black invisible peer-checked/light:visible\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = icons.SunFill(icons.Props{Size: "20", Class: "scale-0 duration-200 peer-checked/light:group-[]/light:scale-100"}).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</label> <input class=\"peer/dark appearance-none absolute\" type=\"radio\" name=\"theme\" value=\"dark\" id=\"theme-dark\" onchange=\"onThemeChange(this)\"> <label for=\"theme-system\" class=\"group/dark absolute flex items-center justify-center w-9 h-9 rounded-full bg-black-950 text-white invisible peer-checked/dark:visible\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = icons.MoonFill(icons.Props{Size: "20", Class: "scale-0 duration-200 peer-checked/dark:group-[]/dark:scale-100"}).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</label></div></div></section>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Var3 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
			templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
			templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
			if !templ_7745c5c3_IsBuffer {
				defer func() {
					templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
					if templ_7745c5c3_Err == nil {
						templ_7745c5c3_Err = templ_7745c5c3_BufErr
					}
				}()
			}
			ctx = templ.InitializeContext(ctx)
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<script>\n\t\t\tlet THEME_STORAGE_KEY = \"iota-theme\";\n\t\t\tlet savedTheme = window.localStorage.getItem(THEME_STORAGE_KEY);\n\t\t\tlet initialTheme = savedTheme ?? \"system\";\n\t\t\tlet root = document.documentElement;\n\t\t\tlet previousTheme = initialTheme;\n\t\t\tlet radioInput = document.getElementById(`theme-${initialTheme}`);\n\t\t\tfunction changeTheme(theme) {\n\t\t\t\troot.classList.remove(previousTheme);\n\t\t\t\tif (!theme) theme = initialTheme;\n\t\t\t\twindow.localStorage.setItem(THEME_STORAGE_KEY, theme);\n\t\t\t\troot.classList.add(theme)\n\t\t\t\tpreviousTheme = theme;\n\t\t\t}\n\t\t\tfunction onThemeChange(input) {\n\t\t\t\tchangeTheme(input.value);\n\t\t\t}\n\t\t\tif (radioInput) {\n\t\t\t\tradioInput.checked = true;\n\t\t\t\tchangeTheme(initialTheme);\n\t\t\t}\n\t\t</script>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			return templ_7745c5c3_Err
		})
		templ_7745c5c3_Err = navbarOnce.Once().Render(templ.WithChildren(ctx, templ_7745c5c3_Var3), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func Index(p *LoginProps) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var4 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var4 == nil {
			templ_7745c5c3_Var4 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Var5 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
			templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
			templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
			if !templ_7745c5c3_IsBuffer {
				defer func() {
					templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
					if templ_7745c5c3_Err == nil {
						templ_7745c5c3_Err = templ_7745c5c3_BufErr
					}
				}()
			}
			ctx = templ.InitializeContext(ctx)
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"flex flex-col h-screen overflow-y-auto\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = Header().Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"flex-1 flex items-center justify-center\"><form class=\"mx-4 max-w-md w-full p-6 md:p-11 flex flex-col gap-4 bg-surface-300 rounded-xl shadow-[0_20px_20px_0px_rgba(0,0,0,0.08),0_0_0_7px_rgba(255,255,255,0.5)]\" method=\"post\"><div class=\"text-center\"><h1 class=\"text-2xl text-100\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var6 string
			templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(p.T("Login.WelcomeBack"))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `modules/core/presentation/templates/pages/login/index.templ`, Line: 82, Col: 33}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</h1><p class=\"mt-2 text-200\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var7 string
			templ_7745c5c3_Var7, templ_7745c5c3_Err = templ.JoinStringErrs(p.T("Login.EnterEmailAndPassword"))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `modules/core/presentation/templates/pages/login/index.templ`, Line: 84, Col: 67}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var7))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</p></div><hr class=\"border border-primary\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if len(p.ErrorMessage) > 0 {
				templ_7745c5c3_Var8 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
					templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
					templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
					if !templ_7745c5c3_IsBuffer {
						defer func() {
							templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
							if templ_7745c5c3_Err == nil {
								templ_7745c5c3_Err = templ_7745c5c3_BufErr
							}
						}()
					}
					ctx = templ.InitializeContext(ctx)
					var templ_7745c5c3_Var9 string
					templ_7745c5c3_Var9, templ_7745c5c3_Err = templ.JoinStringErrs(p.ErrorMessage)
					if templ_7745c5c3_Err != nil {
						return templ.Error{Err: templ_7745c5c3_Err, FileName: `modules/core/presentation/templates/pages/login/index.templ`, Line: 89, Col: 23}
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var9))
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					return templ_7745c5c3_Err
				})
				templ_7745c5c3_Err = alert.Error().Render(templ.WithChildren(ctx, templ_7745c5c3_Var8), templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			templ_7745c5c3_Err = input.Email(&input.Props{
				Label: p.T("Login.Email"),
				Attrs: templ.Attributes{
					"name":  "Email",
					"value": p.Email,
				},
				Error: p.ErrorsMap["Email"],
			}).Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = input.Password(&input.Props{
				Label: p.T("Login.Password"),
				Attrs: templ.Attributes{
					"name": "Password",
				},
				Error: p.ErrorsMap["Password"],
			}).Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Var10 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
				templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
				templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
				if !templ_7745c5c3_IsBuffer {
					defer func() {
						templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
						if templ_7745c5c3_Err == nil {
							templ_7745c5c3_Err = templ_7745c5c3_BufErr
						}
					}()
				}
				ctx = templ.InitializeContext(ctx)
				var templ_7745c5c3_Var11 string
				templ_7745c5c3_Var11, templ_7745c5c3_Err = templ.JoinStringErrs(p.T("Login.Login"))
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `modules/core/presentation/templates/pages/login/index.templ`, Line: 114, Col: 26}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var11))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				return templ_7745c5c3_Err
			})
			templ_7745c5c3_Err = button.Primary(button.Props{
				Size:  button.SizeNormal,
				Class: "justify-center",
				Attrs: templ.Attributes{
					"type": "submit",
				},
			}).Render(templ.WithChildren(ctx, templ_7745c5c3_Var10), templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<p class=\"text-100 text-center\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var12 string
			templ_7745c5c3_Var12, templ_7745c5c3_Err = templ.JoinStringErrs(p.T("Login.OrSignInWith"))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `modules/core/presentation/templates/pages/login/index.templ`, Line: 116, Col: 64}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var12))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</p>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Var13 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
				templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
				templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
				if !templ_7745c5c3_IsBuffer {
					defer func() {
						templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
						if templ_7745c5c3_Err == nil {
							templ_7745c5c3_Err = templ_7745c5c3_BufErr
						}
					}()
				}
				ctx = templ.InitializeContext(ctx)
				templ_7745c5c3_Err = GoogleIcon().Render(ctx, templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				return templ_7745c5c3_Err
			})
			templ_7745c5c3_Err = button.Secondary(button.Props{
				Size:  button.SizeNormal,
				Class: "justify-center",
				Href:  p.GoogleOAuthCodeURL,
				Attrs: templ.Attributes{
					"type": "button",
				},
			}).Render(templ.WithChildren(ctx, templ_7745c5c3_Var13), templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</form></div></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			return templ_7745c5c3_Err
		})
		templ_7745c5c3_Err = layouts.Base(p.Title).Render(templ.WithChildren(ctx, templ_7745c5c3_Var5), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func GoogleIcon() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var14 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var14 == nil {
			templ_7745c5c3_Var14 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"72\" height=\"24\" viewBox=\"0 0 72 24\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><g clip-path=\"url(#clip0_889_9981)\"><path d=\"M30.7174 12.3076C30.7174 15.6389 28.1113 18.0937 24.9131 18.0937C21.7148 18.0937 19.1087 15.6389 19.1087 12.3076C19.1087 8.95279 21.7148 6.52148 24.9131 6.52148C28.1113 6.52148 30.7174 8.95279 30.7174 12.3076ZM28.1765 12.3076C28.1765 10.2258 26.6661 8.80148 24.9131 8.80148C23.16 8.80148 21.6496 10.2258 21.6496 12.3076C21.6496 14.3684 23.16 15.8137 24.9131 15.8137C26.6661 15.8137 28.1765 14.3658 28.1765 12.3076Z\" fill=\"#EA4335\"></path> <path d=\"M43.2391 12.3076C43.2391 15.6389 40.633 18.0937 37.4348 18.0937C34.2365 18.0937 31.6304 15.6389 31.6304 12.3076C31.6304 8.9554 34.2365 6.52148 37.4348 6.52148C40.633 6.52148 43.2391 8.95279 43.2391 12.3076ZM40.6983 12.3076C40.6983 10.2258 39.1878 8.80148 37.4348 8.80148C35.6817 8.80148 34.1713 10.2258 34.1713 12.3076C34.1713 14.3684 35.6817 15.8137 37.4348 15.8137C39.1878 15.8137 40.6983 14.3658 40.6983 12.3076Z\" fill=\"#FBBC05\"></path> <path d=\"M55.2391 6.87105V17.2589C55.2391 21.5319 52.7191 23.2771 49.74 23.2771C46.9357 23.2771 45.2478 21.4015 44.6113 19.8676L46.8235 18.9467C47.2174 19.8884 48.1826 20.9997 49.7374 20.9997C51.6444 20.9997 52.8261 19.8232 52.8261 17.6084V16.7763H52.7374C52.1687 17.478 51.0731 18.091 49.6905 18.091C46.7974 18.091 44.147 15.571 44.147 12.3284C44.147 9.06235 46.7974 6.52148 49.6905 6.52148C51.0705 6.52148 52.1661 7.13453 52.7374 7.8154H52.8261V6.87366H55.2391V6.87105ZM53.0061 12.3284C53.0061 10.291 51.647 8.80148 49.9174 8.80148C48.1644 8.80148 46.6957 10.291 46.6957 12.3284C46.6957 14.345 48.1644 15.8137 49.9174 15.8137C51.647 15.8137 53.0061 14.345 53.0061 12.3284Z\" fill=\"#4285F4\"></path> <path d=\"M59.2174 0.783203V17.7397H56.7391V0.783203H59.2174Z\" fill=\"#34A853\"></path> <path d=\"M68.8748 14.2126L70.847 15.5274C70.2104 16.4691 68.6765 18.0917 66.0261 18.0917C62.7391 18.0917 60.2844 15.5508 60.2844 12.3056C60.2844 8.86475 62.76 6.51953 65.7418 6.51953C68.7444 6.51953 70.2131 8.9091 70.6931 10.2004L70.9565 10.8578L63.2218 14.0613C63.8139 15.2221 64.7348 15.8143 66.0261 15.8143C67.32 15.8143 68.2174 15.1778 68.8748 14.2126ZM62.8044 12.1308L67.9748 9.98388C67.6905 9.26127 66.8348 8.75779 65.8278 8.75779C64.5365 8.75779 62.7391 9.89779 62.8044 12.1308Z\" fill=\"#EA4335\"></path> <path d=\"M9.72784 10.803V8.34826H18C18.0809 8.77609 18.1226 9.28218 18.1226 9.83C18.1226 11.6717 17.6191 13.9491 15.9965 15.5717C14.4183 17.2152 12.4017 18.0917 9.73044 18.0917C4.77914 18.0917 0.615662 14.0587 0.615662 9.10739C0.615662 4.15609 4.77914 0.123047 9.73044 0.123047C12.4696 0.123047 14.4209 1.19783 15.887 2.5987L14.1548 4.33087C13.1035 3.34479 11.6791 2.57783 9.72784 2.57783C6.11218 2.57783 3.28436 5.49174 3.28436 9.10739C3.28436 12.723 6.11218 15.637 9.72784 15.637C12.0731 15.637 13.4087 14.6952 14.2644 13.8396C14.9583 13.1457 15.4148 12.1544 15.5948 10.8004L9.72784 10.803Z\" fill=\"#4285F4\"></path></g> <defs><clipPath id=\"clip0_889_9981\"><rect width=\"70.9565\" height=\"24\" fill=\"white\" transform=\"translate(0.521729)\"></rect></clipPath></defs></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
