# Tiptap Example

Rich editor with [Tiptap](https://tiptap.dev), Tailwind CSS, Headless UI, and Vue 3 as frontend, combined with Go as backend to store images.

> Now the editor can be used in any components (use v-model, [check App.vue](https://github.com/kido1611/tiptap-vuejs-example/blob/main/src/App.vue#L19)).

![example](/images/example.png)

## Feature list

- History (Undo, Redo)
- Heading (1,2,3)
- Text style (Bold, Italic, Underline, Strikethrough)
- List (Ordered, Unordered)
- Link
- Image (+ example with Go as server)
- Blockquote
- Table (Delete table, add column, delete column, add row, delete row, merge or split cell)
- Youtube Video
- Horizontal line

## Preview

[https://kido1611-tiptap-vuejs-example.vercel.app/](https://kido1611-tiptap-vuejs-example.vercel.app/)

## Running Frontend

```shell
cp .env.example .env
npm install
npm run dev
```

## Running Backend

```shell
cd server
go get
go run cmd/api/main.go
```

## Icon

- [Tabler Icon](https://tabler-icons.io/)
- [Material Design Icon](https://materialdesignicons.com/)
