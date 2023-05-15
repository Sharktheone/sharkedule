import {createBrowserRouter, RouteObject} from "react-router-dom"
import {route, routes} from "./routes"
import Kanban from "../pages/task/kanban/kanban"


function makeRoutes() {
    return routes.map(toRouteObject)
}

function toRouteObject(route: route): RouteObject {
    return {
        path: route.path,
        element: route.element,
        children: route.children?.map(toRouteObject)
    }
}


export const router = createBrowserRouter([
    {
        path: "/",
        element: <Kanban/>
    },
    {
        path: "/dashboard",
    },
    {
        path: "/login",
    },
    {
        path: "/register",
    },
    {
        path: "/task",

        children: [
            {
                path: "assigned"
            },
            {
                path: "new"
            },
            {
                path: ":uuid"
            }
        ]
    },
    {
        path: "/calendar",
        children: [
            {
                path: "new"
            },
        ]
    },
    {
        path: "/tickets",
        children: [
            {
                path: "assigned"
            },
            {
                path: "new"
            },
            {
                path: ":uuid"
            }

        ]
    },
    {
        path: "/admin",
        children: [
            {
                path: "users"
            },
            {
                path: "roles"
            },
            {
                path: "permissions"
            },
            {
                path: "settings"
            }
        ]
    }
])