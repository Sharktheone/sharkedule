import {api} from "@/api/api"
import {environment} from "@kanban/types"


export default async function KanbanBoardLoader({params}: {
    params: any
}) {
    const uuid = params.uuid

    return api.get(`/kanban/board/${uuid}`).then((res) => {
        return res.data as environment
    })

}