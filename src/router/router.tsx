import {createBrowserRouter, Link, RouteObject} from "react-router-dom"
import {route, routes} from "./routes"
import Kanban from "../pages/task/kanban/kanban"
import KanbanBoardLoader from "../pages/task/kanban/loader"
import boardsLoader from "../pages/task/boardsloader"
import {lazy} from "react"


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


function Home() {
    return (
        <div>
            <Link to="/kanbanboard">Tasks</Link>
        </div>
    )
}


const KanbanBoards = lazy(() => import("@/pages/task/kanbanboards"))
const BoardsError = lazy(() => import("@/pages/task/boardserror"))
export const router = createBrowserRouter([
    {
        path: "/",
        element: <Home/>,
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
        path: "/kanbanboard",
        element: <KanbanBoards/>,
        loader: boardsLoader,
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
        path: "/kanbanboard/:uuid",
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