package com.xhb.logic.http.packet.test;

import com.xhb.core.packet.HttpPacket;
import com.xhb.core.network.HttpRequestClient;
import com.xhb.logic.http.packet.test.model.*;

public class TestPacket extends HttpPacket<Response> {
	private String name;

	public TestPacket(String name) {
		super(EmptyRequest.instance);
		
		this.name = name;
    }

	@Override
    public HttpRequestClient.Method requestMethod() {
        return HttpRequestClient.Method.GET;
    }

	@Override
    public String requestUri() {
        return "from/" + name;
    }
}
