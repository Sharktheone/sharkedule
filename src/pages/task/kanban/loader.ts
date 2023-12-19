import {api} from "@/api/api"
import {environment} from "@kanban/types"


export default async function KanbanBoardLoader({params}: {
    params: any
}) {
    const workspace = params.workspace
    const board = params.board

    return api.get(`${workspace}/kanban/board/${board}`).then((res) => {
        let env = res.data as environment
        env.workspace = workspace
        return env
    })

}