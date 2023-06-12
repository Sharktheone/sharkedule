import {Dispatch, RefObject, SetStateAction, useState} from "react"
import {notifications} from "@mantine/notifications"
import {api} from "@/api/api"
import {useNavigate} from "react-router-dom"


export class handlers {
    private readonly navigate: ReturnType<typeof useNavigate>
    private readonly setIsAdding: Dispatch<SetStateAction<boolean>>
    private readonly newColRef: RefObject<HTMLInputElement>
    private readonly removeTimeout: number | undefined
    private readonly setRemoveTimeout: Dispatch<SetStateAction<number | undefined>>
    private readonly uuid: string

    constructor(setIsAdding: Dispatch<SetStateAction<boolean>>, newColRef: RefObject<HTMLInputElement>, uuid: string) {
        const [removeTimeout, setRemoveTimeout] = useState<number | undefined>(undefined)

        this.navigate = useNavigate()
        this.setIsAdding = setIsAdding
        this.newColRef = newColRef
        this.removeTimeout = removeTimeout
        this.setRemoveTimeout = setRemoveTimeout
        this.uuid = uuid
    }

    handleNewColumn() {
        this.setIsAdding(true)
    }

    cancelAddColumn() {
        this.setRemoveTimeout(
            setTimeout(() => {
                this.setIsAdding(false)
            }, 100)
        )
    }

    addColumn() {
        if (this.removeTimeout) clearTimeout(this.removeTimeout)

        const name = this.newColRef.current?.value
        if (!name) {
            notifications.show({title: "Error", message: "Column name cannot be empty", color: "red"})
            return
        }

        api.put(`/kanban/board/${this.uuid}/column/new`, {name: name}).then(
            (res) => {
                if (res.status > 300) {
                    notifications.show({title: "Error", message: res.data, color: "red"})
                    console.log(res)
                } else {
                    notifications.show({title: "Success", message: "Column created", color: "green"})
                    if (this.newColRef.current) this.newColRef.current.value = ""

                    this.refresh()
                }
            }).catch(
            (err) => {
                notifications.show({title: "Error", message: err.message, color: "red"})
                console.log(err)
            }
        )
    }

    private refresh() {
        this.navigate("")
    }
}