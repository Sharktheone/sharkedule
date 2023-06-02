import {api} from "@/api/api"


export default async function KanbanBoardLoader({params}: { params: any }) {
    const uuid = params.uuid

    console.log(uuid)

    return api.get(`/kanbanboard/${uuid}`).then((res) => {
        return res.data
    })

}