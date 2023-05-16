import {createBrowserRouter, RouteObject} from "react-router-dom"
import {route, routes} from "./routes"
import Kanban from "../pages/task/kanban/kanban"
import KanbanBoardLoader from "../pages/task/kanban/loader"
import KanbanBoards from "../pages/task/kanbanboards"
import boardsLoader from "../pages/task/boardsloader"
import boardsError from "../pages/task/boardserror"
import BoardsError from "../pages/task/boardserror"


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
        element: <KanbanBoards/>,
        loader:  boardsLoader,
        errorElement: <BoardsError/>,
        children: [
            {
                path: "assigned"
            },
            {
                path: "new"
            },
        ]
    },
    {
        path: "/task/:uuid",
        loader: KanbanBoardLoader,
        element: <Kanban/>,
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