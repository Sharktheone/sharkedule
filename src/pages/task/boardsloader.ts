import {api} from "@/api/api"
import {kanbanBoardType} from "@kanban/types"


export default function boardsLoader() {
    return api.get("/kanban/board/list/names").then((res) => {
        if (res.data == null) {
            return [] as kanbanBoardType[]
        }
        return res.data as kanbanBoardType[]
    })

}