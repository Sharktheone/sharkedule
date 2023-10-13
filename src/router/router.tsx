import {createBrowserRouter, Link} from "react-router-dom"
import KanbanBoardLoader from "@/pages/task/kanban/loader"
import boardsLoader from "@/pages/task/boardsloader"
import {lazy} from "react"



function Home() {
    return (
        <div>
            <Link to="/kanbanboard">Tasks</Link>
        </div>
    )
}


const KanbanBoards = lazy(() => import("@/pages/task/kanbanboards"))
const BoardsError = lazy(() => import("@/pages/task/boardserror"))
const Kanban = lazy(() => import("@/pages/task/kanban/kanban"))
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