import {api} from "@/api/api"
import {environment} from "@kanban/types"


export default async function KanbanBoardLoader({params}: {
    params: any
}) {
    const workspace = params.workspace
    const board = params.board

    return api.get(`${workspace}/kanban/board/${board}`).then((res) => {
        return res.data as environment
    })

}