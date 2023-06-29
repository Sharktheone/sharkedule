import {useContext, useState} from "react"
import {EnvironmentContext} from "@kanban/environment"
import styles from "./styles.module.scss"

export default function TagSelector() {

    const {environment} = useContext(EnvironmentContext)

    const tags = environment.tags

    const [selected, setSelected] = useState<string[]>([])

    return (
        <>
            <div className={styles.avaiableTags}>
                {tags.map((tag) => (
                    <>
                        <label key={tag.uuid} className={styles.tag} style={{
                            backgroundColor: `${tag.color}90`,
                        }}
                        htmlFor={`tag-select-${tag.uuid}`}
                        >
                            {tag.name}
                        </label>
                        <input id={`tag-select-${tag.uuid}`} type="checkbox" className={styles.checkbox} checked={selected.includes(tag.uuid)} onChange={() => {
                            if (selected.includes(tag.name)) {
                                setSelected(selected.filter((t: string) => t !== tag.name))
                            } else {
                                setSelected([...selected, tag.uuid])
                            }
                        }}/>
                    </>
                ))}

            </div>
        </>
    )
}