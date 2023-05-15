import {api} from "../../api/api"
import {kanbanBoardType} from "./kanban/types"


export default function boardsLoader() {
    return api.get("/task/list/names").then((res) => {
        return res.data as kanbanBoardType[]
    })

}