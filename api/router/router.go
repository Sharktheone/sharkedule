package router

import (
	"encoding/json"
	"github.com/Sharktheone/sharkedule/api/element"
	"github.com/Sharktheone/sharkedule/api/workspace"
	"github.com/Sharktheone/sharkedule/web"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

/// # Auth implementation options:
///
/// - handle it at the root level (r.Use(auth))
///		- handle each route in a switch case
///		- Pros/Cons:
///			+ clean at router level
///			- need for each route to be defined in two places
/// - handle it at the route level (r.Get("route", auth, handler))
///		- Pros/Cons:
/// 		- bloated handler definition
///			- harder to see auth when looking at the handler
///			+ you can see every route's auth in one place
/// 		+ handler is clean from any auth code
/// - handle it at the handler level (handler(c *fiber.Ctx, auth bool))
///		- Pros/Cons:
///			- handler has auth code
///			+ handler definition is clean
/// - pass auth "decision" to user struct
/// 	- __Handler calls user.<func>() => user checks if it is allowed and returns an error if not__
///		- Pros/Cons:
///			- handler has more code
///			+ handler definition is clean
///		- Handler calls user.<func>(workspace) => user checks if it is allowed and calls workspace.<func>() if it is
///			- Pros/Cons:
///				+ clean handler
///				- (user has more code)
/// - something like user.workspace(workspace).<func>()
///
/// - Handle it with middleware (ExtractUser)
///		- Pros/Cons:
///			+ we need to do it anyway
///			+ handler has not more code than a normal handler

func Start() {
	r := fiber.New(fiber.Config{
		JSONDecoder: json.Unmarshal,
		JSONEncoder: json.Marshal,
	})

	r.Use(cors.New())

	r.Use(logger.New())
	r.Use(recover.New())

	api := r.Group("api") // /api

	workspaces := api.Group("workspace") // /api/workspace
	{
		workspaces.Get("", workspace.List)            // GET /api/workspace
		workspaces.Post("", workspace.ListWithFields) // POST /api/workspace
		workspaces.Get("info", workspace.Info)        // GET /api/workspace/info
	}

	ws := api.Group(":workspace") // /api/:workspace
	{
		ws.Get("", workspace.Info)      // GET /api/:workspace
		ws.Delete("", workspace.Delete) // DELETE /api/:workspace

		elem := ws.Group(":element") // /api/:workspace/:element
		{
			elem.Get("", element.Info)      // GET /api/:workspace/:element
			elem.Put("", element.Create)    // PUT /api/:workspace/:element
			elem.Patch("", element.Update)  // POST /api/:workspace/:element
			elem.Delete("", element.Delete) // DELETE /api/:workspace/:element

			elem.Get("attachments", element.Attachments)           // GET /api/:workspace/:element/attachments
			elem.Put("attach", element.Attach)                     // POST /api/:workspace/:element/attach
			elem.Delete("detach", element.Detach)                  // DELETE /api/:workspace/:element/detach
			elem.Patch("move", element.Move)                       // POST /api/:workspace/:element/move
			elem.Patch("move_element", element.MoveElement)        // POST /api/:workspace/:element/move_element
			elem.Patch("copy", element.Copy)                       // POST /api/:workspace/:element/copy
			elem.Patch("copy_element", element.CopyElement)        // POST /api/:workspace/:element/copy_element
			elem.Get("types", element.GetType)                     // GET /api/:workspace/:element/types
			elem.Patch("types", element.UpdateType)                // PATCH /api/:workspace/:element/types
			elem.Get("subelements", element.List)                  // GET /api/:workspace/:element/subelements
			elem.Get("subelements/:type", element.ListType)        // GET /api/:workspace/:element/subelements/:type
			elem.Get("rec_subelements", element.RecList)           // GET /api/:workspace/:element/rec_subelements
			elem.Get("rec_subelements/:type", element.RecListType) // GET /api/:workspace/:element/rec_subelements/:type

			fields := elem.Group("fields") // /api/:workspace/:uuid/fields
			{
				fields.Get("", element.FieldList)           // GET /api/:workspace/:element/fields
				fields.Get(":field", element.FieldInfo)     // GET /api/:workspace/:element/fields/:field
				fields.Patch(":field", element.FieldUpdate) // PATCH /api/:workspace/:element/fields/:field
			}

		}
	}

	web.Serve(r)

	if err := r.Listen(":5639"); err != nil {
		panic(err)
	}

}
