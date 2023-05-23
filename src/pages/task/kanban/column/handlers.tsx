import styles from "./styles.module.scss"
import {api} from "../../../../api/api"
import {notifications} from "@mantine/notifications"
import {IconX} from "@tabler/icons-react"
import {Dispatch, RefObject, SetStateAction, useEffect, useRef, useState} from "react"
import {kanbanBoardType} from "../types"
import {useNavigate} from "react-router-dom"
import {ghostElementType, ghostType} from "../ghost"

export class handlers {
    readonly editable: boolean
    readonly tasksRef: RefObject<HTMLDivElement>
    readonly nameRef: RefObject<HTMLTextAreaElement>
    readonly ghostElement: ghostElementType | undefined
    readonly isAdding: boolean
    private readonly setEditable: Dispatch<SetStateAction<boolean>>
    private readonly uuid: string
    private readonly board: kanbanBoardType
    private readonly setBoard: Dispatch<SetStateAction<kanbanBoardType>>
    private readonly navigate: ReturnType<typeof useNavigate>
    private readonly setGhostElement: Dispatch<SetStateAction<ghostElementType | undefined>>
    private readonly setIsAdding: Dispatch<SetStateAction<boolean>>
    private readonly removeTimeout: number | undefined
    private readonly setRemoveTimeout: Dispatch<SetStateAction<number | undefined>>
    private readonly ghost: ghostType | undefined


    constructor(uuid: string, setBoard: Dispatch<SetStateAction<kanbanBoardType>>, board: kanbanBoardType, ghost: ghostType | undefined) {
        const [ghostElement, setGhostElement] = useState<ghostElementType | undefined>()
        const [removeTimeout, setRemoveTimeout] = useState<number | undefined>(undefined)
        const [editable, setEditable] = useState(false)
        const [isAdding, setIsAdding] = useState(false)

        const nameRef = useRef<HTMLTextAreaElement>(null)
        const tasksRef = useRef<HTMLDivElement>(null)

        this.setGhostElement = setGhostElement
        this.setRemoveTimeout = setRemoveTimeout
        this.setEditable = setEditable
        this.setIsAdding = setIsAdding
        this.editable = editable
        this.ghostElement = ghostElement
        this.removeTimeout = removeTimeout
        this.isAdding = isAdding
        this.uuid = uuid
        this.board = board
        this.setBoard = setBoard
        this.navigate = useNavigate()
        this.tasksRef = tasksRef
        this.nameRef = nameRef
        this.ghost = ghost
    }



    renameColumn(uuid: string, name: string) {
        let newBoard = {...this.board}
        newBoard.columns?.forEach((column) => {
            if (column.uuid === uuid) {
                column.name = name
                return
            }
        })
        this.setBoard(newBoard)
    }



    editText() {
        this.setEditable(true)
    }

    handleBlur(e: any) {
        this.setEditable(false)
        this.renameColumn(this.uuid, e.target.innerText)
    }


    renameTask(uuid: string, name: string) {
        let newBoard = {...this.board}
        newBoard.columns?.forEach((column) => {
            column.tasks?.forEach((task) => {
                if (task.uuid === uuid) {
                    task.name = name
                    return
                }
            })
        })
        this.setBoard(newBoard)
    }

    handleDelete() {
        api.delete(`/kanbanboard/${this.board.uuid}/column/${this.uuid}/delete`).then(
            (res) => {
                if (res.status > 300) {
                    notifications.show({title: "Error", message: "res.data", color: "red", icon: <IconX/>})
                } else {
                    notifications.show({title: "Success", message: "Deleted Column", color: "green"})
                    this.refresh()

                }
            }
        )
    }

    checkGhost() {

        useEffect(() => {
            if (!this.ghost) {
                this.setGhostElement(undefined)
                return
            }
            if (this.ghost.hoveredColumnID !== this.uuid) {
                this.setGhostElement(undefined)
                return
            }

            let offset = 0

            let tasks = [].slice.call(this.tasksRef.current?.children) as HTMLDivElement[]

            console.log(tasks.filter(task => !task.className.includes(styles.ghost)))

            tasks.filter(task => !task.className.includes(styles.ghost)).forEach((task, index) => {
                if (index < this.ghost!.index) {
                    offset += task.getBoundingClientRect().height
                }
            })

            const ghostElement = {
                height: this.ghost.height + "px",
                offsetTop: offset + "px",
            }

            console.log(ghostElement)

            this.setGhostElement(ghostElement)
        }, [this.ghost])

    }
    handleNewTask() {
        this.setIsAdding(true)
    }

    checkAdding() {
        useEffect(() => {
            if (this.isAdding) {
                this.nameRef.current?.focus()
            }
        }, [this.isAdding])
    }


    addTask() {
        if (this.removeTimeout) clearTimeout(this.removeTimeout)

        let name: string
        if (this.nameRef.current?.value) {
            name = this.nameRef.current?.value
        } else {
            notifications.show({title: "Error", message: "Task name cannot be empty", color: "red", icon: <IconX/>})
            return
        }

        api.put(`/kanbanboard/${this.board.uuid}/column/${this.uuid}/task/new`, {name: name}).then(
            (res) => {
                if (res.status > 300) {
                    notifications.show({title: "Error", message: "res.data", color: "red", icon: <IconX/>})
                } else {
                    this.renameTask(res.data.uuid, name)
                    if (this.nameRef.current) {
                        this.nameRef.current.value = ""
                        this.nameRef.current.focus()
                    }
                    this.refresh()
                }

            }).catch(e => {
            notifications.show({title: "Error", message: e.message, color: "red", icon: <IconX/>})
        })
    }


    removeIsAdding() {
        if (this.removeTimeout) clearTimeout(this.removeTimeout)
        this.setRemoveTimeout(setTimeout(() => {
            this.setIsAdding(false)
        }, 100))
    }

    closeIsAdding() {
        if (this.removeTimeout) clearTimeout(this.removeTimeout)
        this.setIsAdding(false)
    }

    private refresh() {
        this.navigate("")
    }
}