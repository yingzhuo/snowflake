### Java客户端

由于接口简单，不单独封装Java客户端。请参考以下代码。

```java
package com.github.yingzhuo;

import org.junit.Test;
import org.springframework.core.ParameterizedTypeReference;
import org.springframework.http.HttpMethod;
import org.springframework.http.ResponseEntity;
import org.springframework.web.client.RestTemplate;

import java.util.Collections;
import java.util.List;
import java.util.Objects;

public class TestCases {

    private static final RestTemplate REST_TEMPLATE = new RestTemplate();

    @Test
    public void test() {

        final ResponseEntity<List<Long>> response = REST_TEMPLATE.exchange(
                "http://<host>:<port>/id?n={n}",
                HttpMethod.GET,
                null,
                new ParameterizedTypeReference<List<Long>>() {
                },
                Collections.singletonMap("n", 10)
        );

        for (Long id : Objects.requireNonNull(response.getBody())) {
            System.out.println(id);
        }
    }

}
```