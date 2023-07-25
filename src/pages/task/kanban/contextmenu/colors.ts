import {createStyles} from "@mantine/core"


export const useColors = createStyles((theme) => ({
        contextMenu: {
            backgroundColor: theme.colors.dark[4],
            border: `1px solid ${
                theme.colorScheme === "dark" ? theme.colors.dark[3] : theme.colors.gray[2]
            }`,
        }
    }
))