import {api} from "@/api/api"
import {NameList} from "@kanban/types"


export default function boardsLoader() {
    return api.get("/workspace").then((res) => {
        if (res.data == null) {
            return [] as NameList[]
        }
        return res.data as NameList[]
    })

}