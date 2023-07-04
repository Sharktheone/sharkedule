import {useContext, useEffect, useState} from "react"
import {EnvironmentContext} from "@kanban/environment"
import styles from "./styles.module.scss"
import {useColors} from "@/components/kanban/tags/selector/colors"
import {useClickOutside} from "@mantine/hooks"
import {IconCirclePlus} from "@tabler/icons-react"

type Props = {
    onChange: (tags: string[]) => void
    selected?: string[]
}
export default function TagSelector({onChange, selected}: Props) {
    const {environment} = useContext(EnvironmentContext)

    const tags = environment?.tags

    const [newSelected, setNewSelected] = useState<string[]>(selected ?? [])
    const [firstRender, setFirstRender] = useState<boolean>(true)
    const [opened, setOpened] = useState<boolean>(false)
    const [selectedRef, setSelectedRef] = useState<HTMLDivElement | null>(null)

    const [popoverRef, setPopoverRef] = useState<HTMLDivElement | null>(null)
    useClickOutside(() => setOpened(false), null, [selectedRef, popoverRef])

    const {classes, cx} = useColors()

    useEffect(() => {
        setNewSelected(selected ?? [])
    }, [selected])

    useEffect(() => {
        if (firstRender) {
            setFirstRender(false)
            return
        }
        onChange(newSelected)
    }, [newSelected])

    function checked(uuid: string) {
        return newSelected.includes(uuid)
    }

    function handleChange(uuid: string) {
        if (newSelected.includes(uuid)) {
            setNewSelected(newSelected.filter((t: string) => t !== uuid))
        } else {
            setNewSelected([...newSelected, uuid])
        }
    }

    function deleteTag(uuid: string) {
        if (newSelected.includes(uuid)) {
            setNewSelected(newSelected.filter((t: string) => t !== uuid))
        } else {
            setNewSelected([...newSelected, uuid])
        }
    }

    function open() {
        setOpened(true)
    }

    //@ts-ignore
    function fakeContextMenu(e: MouseEvent<HTMLDivElement, MouseEvent>, uuid: string) {
        e.preventDefault()
        deleteTag(uuid)
    }

    return (
        <div className={styles.selector}>
            <div ref={setSelectedRef} className={styles.selected}>
                {newSelected?.map((uuid) => {
                    let tag = tags?.find((tag) => tag.uuid === uuid)
                    if (!tags) {
                        return null
                    }
                    return (
                        <div key={uuid}
                             onClick={open}
                             onContextMenu={e => fakeContextMenu(e, uuid)}
                             style={{
                                 backgroundColor: `${tag?.color}90`,
                             }}>
                            {tag?.name}
                        </div>
                    )
                })}
                <button onClick={open}>
                    <IconCirclePlus/>
                </button>
            </div>
            {opened ? <div className={`${styles.availableTags} ${cx(classes.availableTags)}`}
                           ref={setPopoverRef}
                >
                    {tags?.map((tag) => (
                        <>
                            <label key={tag.uuid} className={styles.tag}
                                   htmlFor={`tag-select-${tag.uuid}`}
                            >
                                <input id={`tag-select-${tag.uuid}`} type="checkbox" className={styles.checkbox}
                                       checked={checked(tag.uuid)}
                                       onChange={() => handleChange(tag.uuid)}/>

                                <span className={styles.name} style={{
                                    backgroundColor: `${tag.color}90`,
                                }}>
                                {tag.name}
                            </span>
                            </label>
                        </>
                    ))}
                </div>
                : null
            }
        </div>
    )
}