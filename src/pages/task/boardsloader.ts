import {api} from "@/api/api"
import {WorkspaceList} from "@kanban/types"


export default function boardsLoader() {
    return api.get("/workspace/info").then((res) => {
        if (res.data == null) {
            return [] as WorkspaceList[]
        }
        return res.data as WorkspaceList[]
    })

}