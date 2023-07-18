using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class Enemy : MonoBehaviour
{
    // Start is called before the first frame update

    public float speed = 5;
    Vector3 dir;

    // ���� ���� �ּ�
    public GameObject explosionFactory;

    void Start()
    {

        
    }
    void OnEnable()
    {

        int randValue = Random.Range(0, 10);
        if (randValue < 3)
        {
            // �÷��̾� ����
            GameObject target = GameObject.Find("Player");
            dir = target.transform.position - transform.position;
            dir.Normalize();
        }
        else
        {
            // �Ʒ� ����
            dir = Vector3.down;
        }
    }
    // Update is called once per frame
    void Update()
    {

        transform.position += dir * speed * Time.deltaTime;
    }

    // �浹 ���� ��
    private void OnCollisionEnter(Collision other)
    {

        // ���� ����Ʈ�� �߻���Ŵ
        GameObject explosion = Instantiate(explosionFactory);

        // ���� ȿ���� ��ġ��Ŵ
        explosion.transform.position = transform.position;


        // �ε����� ��Ȱ��ȭ���Ѽ� �ٽ� źâ�� �ִ´�

        // �ε��� ����� �̸��� Bullet�̶��
        if (other.gameObject.name.Contains("Bullet"))
        {
            // �ε��� ��ü ��Ȱ��ȭ
            other.gameObject.SetActive(false);
        }

        // �ƴ϶�� ����
        else
        {
            Destroy(other.gameObject);
        }

        // Destroy(gameObject);
        gameObject.SetActive(false); // Ǯ�� �ڿ� �ݳ�

        PlayerFire player = GameObject.Find("Player").GetComponent<PlayerFire>();
        player.bulletObjectPool.Add(other.gameObject);

        // ���� ���� ������ ���� ǥ��

        ScoreManager.Instance.Score++;

    }



}
