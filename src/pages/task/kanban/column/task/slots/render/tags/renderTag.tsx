import styles from "@kanban/column/task/slots/render/tags/styles.module.scss"
import {getTag} from "@/pages/task/utils/tag"
import {Tag} from "@kanban/types"


type Props = {
    tags: string[]
}

export default function RenderTags({tags}: Props) {
    return (
        <div className={styles.tags}>
            {tags?.map((tag) => (
                <RenderTag key={tag} uuid={tag}/>
            ))}
        </div>
    )
}

type RenderTagProps = {
    uuid: string
}

export function RenderTag({uuid}: RenderTagProps) {
    let tag: Tag = getTag(uuid) ?? {name: "Error, Tag not found", color: "#000000"} as Tag

    return (
        <div className={styles.tag} style={{backgroundColor: `${tag.color}80`, borderColor: tag.color}}>
            {tag.name}
        </div>
    )
}