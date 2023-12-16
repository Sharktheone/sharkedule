import {createBrowserRouter} from "react-router-dom"
import KanbanBoardLoader from "@/pages/task/kanban/loader"
import boardsLoader from "@/pages/task/boardsloader"
import {lazy} from "react"
import Home from "@/pages/home/home"
import Login from "@/pages/login/login"


const Kanban = lazy(() => import("@/pages/task/kanban"))
const KanbanError = lazy(() => import("@/pages/task/kanbaberror"))
const Board = lazy(() => import("@/pages/task/kanban/kanban"))
export const router = createBrowserRouter([
    {
        path: "/",
        element: <Home/>,
    },
    // {
    //     path: "/dashboard",
    // },
    {
        path: "/login",
        element: <Login/>,
    },
    {
        path: "/register",
    },
    {
        path: "/dashboard",
        element: <Kanban/>,
        loader: boardsLoader,
        errorElement: <KanbanError/>,
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
        path: "board/:workspace/:board",
        loader: KanbanBoardLoader,
        element: <Board/>,
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