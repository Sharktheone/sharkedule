import {Dispatch, RefObject, SetStateAction, useContext, useState} from "react"
import {notifications} from "@mantine/notifications"
import {api} from "@/api/api"
import {useNavigate} from "react-router-dom"
import {EnvironmentContext} from "@kanban/environment"


export class handlers {
    private readonly navigate: ReturnType<typeof useNavigate>
    private readonly setIsAdding: Dispatch<SetStateAction<boolean>>
    private readonly newColRef: RefObject<HTMLInputElement>
    private readonly removeTimeout: number | undefined
    private readonly setRemoveTimeout: Dispatch<SetStateAction<number | undefined>>
    private readonly uuid: string
    private readonly workspace: string

    constructor(setIsAdding: Dispatch<SetStateAction<boolean>>, newColRef: RefObject<HTMLInputElement>, uuid: string, workspace: string) {
        const [removeTimeout, setRemoveTimeout] = useState<number | undefined>(undefined)

        this.navigate = useNavigate()
        this.setIsAdding = setIsAdding
        this.newColRef = newColRef
        this.removeTimeout = removeTimeout
        this.setRemoveTimeout = setRemoveTimeout
        this.uuid = uuid
        this.workspace = workspace
    }

    handleNewColumn() {
        this.setIsAdding(true)
    }

    cancelAddColumn() {
        this.setRemoveTimeout(
            //@ts-ignore
            setTimeout(() => {
                this.setIsAdding(false)
            }, 500)
        )
    }

    addColumn() {
        if (this.removeTimeout) clearTimeout(this.removeTimeout)

        const name = this.newColRef.current?.value
        if (!name) {
            notifications.show({title: "Error", message: "Column name cannot be empty", color: "red"})
            return
        }

        api.put(`/${this.workspace}/kanban/board/${this.uuid}/column/new`, {name: name}).then(
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
        this.newColRef.current?.focus()

    }

    private refresh() {
        this.navigate("")
    }
}