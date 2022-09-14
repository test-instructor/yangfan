import * as monaco from 'monaco-editor'
import { language as pythonLanguage } from 'monaco-editor/esm/vs/basic-languages/python/python.js';

const pythonCompletion = monaco.languages.registerCompletionItemProvider('python', {
    provideCompletionItems: function () {
        let suggestions = [];
        pythonLanguage.keywords.forEach(item => {
            suggestions.push({
                label: item,
                kind: monaco.languages.CompletionItemKind.Keyword,
                insertText: item
            });
        })
        return {
            suggestions:suggestions
        };
    },
});


export {
    pythonCompletion
};
