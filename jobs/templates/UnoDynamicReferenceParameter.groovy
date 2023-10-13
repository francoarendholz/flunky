            {% autoescape off %}[$class: 'DynamicReferenceParameter',
                name: '{{ name }}',
                description: '{{ description }}',
                referencedParameters: '{{ referencedParameters }}',
                choiceType: '{{ choiceType }}',
                omitValueField: {{ omitValueField }},
                script: [
                    $class: '{{ scriptClass }}', 
                    fallbackScript: [
                        classpath: [{{ fallbackClassPath }}], 
                        sandbox: {{ fallbackSandbox }}, 
                        script: 
                            '''{{ fallbackScript }}'''
                    ], 
                    script: [
                        classpath: [{{ classPath }}], 
                        sandbox: {{ scriptSandbox }}, 
                        script: 
                            '''{{ script }}'''
                    ]
                ]
            ]{% endautoescape %}